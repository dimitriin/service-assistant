package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dimitriin/service-assistant/pkg/protocol"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
	"github.com/golang/protobuf/proto"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	defer func() {
		_ = logger.Sync()
	}()

	conn, err := net.Dial("unixgram", "/var/run/service-assistant/service-assistant.sock")

	if err != nil {
		log.Fatalf("unable to connect to assistant", zap.Error(err))
	}

	defer func() {
		_ = conn.Close()
	}()

	errCh := make(chan error, 1)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	r := mux.NewRouter()
	r.Path("/metrics").Handler(promhttp.Handler())
	r.Path("/healthzBit").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, protocol.HealthzBitType)

		bit := &payload.HealthBit{
			Ttl: 120,
		}

		data, _ := proto.Marshal(bit)

		n, err := fmt.Fprintf(conn, "%s%s", string(buf), string(data))

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(fmt.Sprintf(`{"n":%d,"err":"%s"}`, n, err)))
		} else {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte(fmt.Sprintf(`{"n":%d,"err":null}`, n)))
		}
	})
	r.Path("/readyzBit").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, protocol.ReadyzBitType)

		bit := &payload.ReadyBit{
			Ttl: 120,
		}

		data, _ := proto.Marshal(bit)

		n, err := fmt.Fprintf(conn, "%s%s", string(buf), string(data))

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(fmt.Sprintf(`{"n":%d,"err":"%s"}`, n, err)))
		} else {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte(fmt.Sprintf(`{"n":%d,"err":null}`, n)))
		}
	})

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		errCh <- httpServer.ListenAndServe()
	}()

	go func() {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, protocol.ReadyzBitType)

		bit := &payload.ReadyBit{
			Ttl: 120,
		}

		data, _ := proto.Marshal(bit)

		if _, err := fmt.Fprintf(conn, "%s%s", string(buf), string(data)); err != nil {
			log.Fatalf("unable to send initial ready bit", zap.Error(err))
		}
	}()

	go func() {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, protocol.HealthzBitType)

		bit := &payload.HealthBit{
			Ttl: 120,
		}

		data, _ := proto.Marshal(bit)

		if _, err := fmt.Fprintf(conn, "%s%s", string(buf), string(data)); err != nil {
			log.Fatalf("unable to send initial health bit", zap.Error(err))
		}
	}()

	ctx := context.Background()

	select {
	case err := <-errCh:
		log.Error("service component stopped unexpectedly with error", zap.Error(err))
	case <-sigCh:
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Error("server stopped with error", zap.Error(err))
		}
	}
}
