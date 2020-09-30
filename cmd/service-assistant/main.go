package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/dimitriin/service-assistant/pkg/config"
	"github.com/dimitriin/service-assistant/pkg/handler"
	"github.com/dimitriin/service-assistant/pkg/metrics"
	"github.com/dimitriin/service-assistant/pkg/protocol"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rdz atomic.Value
var hlz atomic.Value

func main() {
	logger, _ := zap.NewProduction()

	defer func() {
		_ = logger.Sync()
	}()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("service.http.address", ":8181")
	viper.SetDefault("service.udp.address", ":8282")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("read config file error", zap.Error(err))
	}

	logger.WithOptions()

	cfg := config.Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatal("unmarshal config error", zap.Error(err))
	}

	validator := config.NewValidator()

	if err := validator.Validate(cfg); err != nil {
		logger.Fatal("validate config error", zap.Error(err))
	}

	registry := metrics.NewRegistry(cfg.Metrics)

	if err := registry.Register(); err != nil {
		logger.Fatal("metrics register error", zap.Error(err))
	}

	pc, err := net.ListenPacket("udp", cfg.Service.UDP.Address)

	if err != nil {
		logger.Fatal("listen error", zap.Error(err))
	}

	processor := protocol.NewPacketStreamProcessor(
		pc,
		protocol.NewDecoder(),
		handler.NewAggregateCmdHandler(
			handler.NewCounterIncHandler(registry),
			handler.NewCounterAddHandler(registry),
		),
		zap.NewNop(),
	)

	errCh := make(chan error, 2)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	go func() {
		errCh <- processor.Process()
	}()

	rdz.Store(true)
	hlz.Store(true)

	r := mux.NewRouter()
	r.Path("/metrics").Handler(promhttp.Handler())
	r.Path("/readyz").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		v := rdz.Load()
		ready, _ := v.(bool)

		if ready {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte(`{"ready":true}`))
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(`{"ready":false}`))
		}
	})
	r.Path("/healthz").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		v := hlz.Load()
		health, _ := v.(bool)

		if health {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte(`{"health":true}`))
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(`{"health":false}`))
		}
	})
	r.Path("/unhealthz").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		hlz.Store(false)
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"health":false}`))
	})
	r.Path("/unreadyz").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		rdz.Store(false)
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"ready":false}`))
	})

	httpServer := &http.Server{
		Addr:    cfg.Service.HTTP.Address,
		Handler: r,
	}

	go func() {
		errCh <- httpServer.ListenAndServe()
	}()

	ctx := context.Background()

	select {
	case err := <-errCh:
		log.Error("service component stopped unexpectedly with error", zap.Error(err))
	case <-sigCh:
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Error("http server stopped with error", zap.Error(err))
		}

		if err := pc.Close(); err != nil {
			log.Error("server stopped with error", zap.Error(err))
		}
	}
}
