package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type CounterIncHandler struct {
	registry *Registry
}

func NewCounterIncHandler(registry *Registry) *CounterIncHandler {
	return &CounterIncHandler{registry: registry}
}

func (h *CounterIncHandler) Handle(value interface{}) error {
	packetIncCmd, ok := value.(*payload.Packet_CounterIncCmd)

	if !ok {
		return errors.New("unexpected value for counter inc handler")
	}

	incCmd := packetIncCmd.CounterIncCmd

	counter, err := h.registry.GetCounter(incCmd.Name)

	if err != nil {
		return err
	}

	m, err := counter.GetMetricWith(incCmd.Labels)

	if err != nil {
		return err
	}

	m.Inc()

	return nil
}
