package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"

	"github.com/dimitriin/service-assistant/pkg/probe"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/dimitriin/service-assistant/pkg/config"
	"github.com/dimitriin/service-assistant/pkg/metrics"
	"github.com/dimitriin/service-assistant/pkg/protocol"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	defer func() {
		_ = logger.Sync()
	}()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.service-assistant/config")
	viper.AddConfigPath("/etc/service-assistant/config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("service.http.address", ":8181")
	viper.SetDefault("service.conn.network", "unixgram")
	viper.SetDefault("service.conn.address", "/var/run/service-assistant/service-assistant.sock")

	if err := viper.ReadInConfig(); err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)

		if !ok {
			logger.Fatal("read config file error", zap.Error(err))
		}

		logger.Warn("configuration file not found", zap.Error(err))
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

	pc, err := net.ListenPacket(cfg.Service.Conn.Network, cfg.Service.Conn.Address)

	if err != nil {
		logger.Fatal("listen error", zap.Error(err))
	}

	rdzHandler := probe.NewHandler()
	rdzHandler.StartTimeBit()

	hlzHandler := probe.NewHandler()
	hlzHandler.StartTimeBit()

	processor := protocol.NewPacketStreamProcessor(
		pc,
		protocol.NewPacketHandler(map[string]protocol.HandlerInterface{
			reflect.TypeOf(&payload.Packet_ReadyBit{}).String():                 rdzHandler,
			reflect.TypeOf(&payload.Packet_HealthBit{}).String():                hlzHandler,
			reflect.TypeOf(&payload.Packet_CounterRegisterCmd{}).String():       metrics.NewCounterRegisterHandler(registry),
			reflect.TypeOf(&payload.Packet_CounterIncCmd{}).String():            metrics.NewCounterIncHandler(registry),
			reflect.TypeOf(&payload.Packet_CounterAddCmd{}).String():            metrics.NewCounterAddHandler(registry),
			reflect.TypeOf(&payload.Packet_HistogramRegisterCmd{}).String():     metrics.NewHistogramRegisterHandler(registry),
			reflect.TypeOf(&payload.Packet_HistogramObserveCmd{}).String():      metrics.NewHistogramObserveHandler(registry),
			reflect.TypeOf(&payload.Packet_GaugeRegisterCmd{}).String():         metrics.NewGaugeRegisterHandler(registry),
			reflect.TypeOf(&payload.Packet_GaugeSetToCurrentTimeCmd{}).String(): metrics.NewGaugeSetToCurrentTimeHandler(registry),
			reflect.TypeOf(&payload.Packet_GaugeAddCmd{}).String():              metrics.NewGaugeAddHandler(registry),
			reflect.TypeOf(&payload.Packet_GaugeDecCmd{}).String():              metrics.NewGaugeDecHandler(registry),
			reflect.TypeOf(&payload.Packet_GaugeIncCmd{}).String():              metrics.NewGaugeIncHandler(registry),
			reflect.TypeOf(&payload.Packet_GaugeSetCmd{}).String():              metrics.NewGaugeSetHandler(registry),
			reflect.TypeOf(&payload.Packet_GaugeSubCmd{}).String():              metrics.NewGaugeSubHandler(registry),
		}),
		logger,
	)

	errCh := make(chan error, 2)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	go func() {
		errCh <- processor.Process()
	}()

	r := mux.NewRouter()
	r.Path("/metrics").Handler(promhttp.Handler())
	r.Path("/readyz").Handler(rdzHandler)
	r.Path("/healthz").Handler(hlzHandler)

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
			log.Error("packet conn error", zap.Error(err))
		}

		if err := syscall.Unlink(cfg.Service.Conn.Address); err != nil {
			log.Error("unlink socket error", zap.Error(err))
		}
	}
}
