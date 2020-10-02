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
	incCmd, ok := value.(*payload.CounterAddCmd)

	if !ok {
		return errors.New("unexpected value for counter inc handler")
	}

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
