package handler

import (
	"github.com/dimitriin/service-assistant/pkg/metrics"
	"github.com/dimitriin/service-assistant/pkg/protocol/cmd"
)

type CounterIncHandler struct {
	registry *metrics.Registry
}

func NewCounterIncHandler(registry *metrics.Registry) *CounterIncHandler {
	return &CounterIncHandler{registry: registry}
}

func (h *CounterIncHandler) handle(cmd *cmd.CounterIncCmd) error {
	counter, err := h.registry.GetCounter(cmd.Name)

	if err != nil {
		return err
	}

	m, err := counter.GetMetricWith(cmd.Labels)

	if err != nil {
		return err
	}

	m.Inc()

	return nil
}
