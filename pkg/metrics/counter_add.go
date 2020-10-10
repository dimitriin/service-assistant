package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type CounterAddHandler struct {
	registry *Registry
}

func NewCounterAddHandler(registry *Registry) *CounterAddHandler {
	return &CounterAddHandler{registry: registry}
}

func (h *CounterAddHandler) Handle(value interface{}) error {
	packetAddCmd, ok := value.(*payload.Packet_CounterAddCmd)

	if !ok {
		return errors.New("unexpected value for counter add handler")
	}

	addCmd := packetAddCmd.CounterAddCmd

	counter, err := h.registry.GetCounter(addCmd.Name)

	if err != nil {
		return err
	}

	m, err := counter.GetMetricWith(addCmd.Labels)

	if err != nil {
		return err
	}

	m.Add(addCmd.Value)

	return nil
}
