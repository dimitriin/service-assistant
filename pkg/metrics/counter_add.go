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
	addCmd, ok := value.(*payload.CounterAddCmd)

	if !ok {
		return errors.New("unexpected value for counter add handler")
	}

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
