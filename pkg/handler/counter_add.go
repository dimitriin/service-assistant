package handler

import (
	"github.com/dimitriin/service-assistant/pkg/metrics"
	"github.com/dimitriin/service-assistant/pkg/protocol/cmd"
)

type CounterAddHandler struct {
	registry *metrics.Registry
}

func NewCounterAddHandler(registry *metrics.Registry) *CounterAddHandler {
	return &CounterAddHandler{registry: registry}
}

func (h *CounterAddHandler) handle(cmd *cmd.CounterAddCmd) error {
	counter, err := h.registry.GetCounter(cmd.Name)

	if err != nil {
		return err
	}

	m, err := counter.GetMetricWith(cmd.Labels)

	if err != nil {
		return err
	}

	m.Add(float64(cmd.Value))

	return nil
}
