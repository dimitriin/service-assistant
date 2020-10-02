package probe

import (
	"errors"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type Handler struct {
	probeValue *atomic.Value
	bitChan    chan uint64
}

func NewHandler() *Handler {
	probeValue := &atomic.Value{}
	probeValue.Store(false)

	h := &Handler{
		probeValue: probeValue,
		bitChan:    make(chan uint64, 1024),
	}

	return h
}

func (h *Handler) StartTimeBit() {
	go func() {
		timer := time.NewTimer(time.Second * 0)

		for {
			select {
			case <-timer.C:
				h.probeValue.Store(false)
				timer.Stop()
			case ttl := <-h.bitChan:
				h.probeValue.Store(true)
				if ttl == 0 {
					timer.Stop()
				} else {
					timer.Reset(time.Second * time.Duration(ttl))
				}
			}
		}
	}()
}

func (h *Handler) Handle(value interface{}) error {
	switch value.(type) {
	case *payload.ReadyBit:
		bit := value.(*payload.ReadyBit)
		h.bitChan <- bit.Ttl
	case *payload.HealthBit:
		bit := value.(*payload.HealthBit)
		h.bitChan <- bit.Ttl
	default:
		return errors.New("unexpected probe handler value")
	}

	return nil
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	probeValue := h.probeValue.Load().(bool)

	if probeValue {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"status":"ok"}`))
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(`{"status":"failed"}`))
	}
}
