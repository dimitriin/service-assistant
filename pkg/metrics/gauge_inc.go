package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type GaugeIncHandler struct {
	registry *Registry
}

func NewGaugeIncHandler(registry *Registry) *GaugeIncHandler {
	return &GaugeIncHandler{registry: registry}
}

func (h *GaugeIncHandler) Handle(value interface{}) error {
	incCmd, ok := value.(*payload.GaugeIncCmd)

	if !ok {
		return errors.New("unexpected value for gauge inc handler")
	}

	counter, err := h.registry.GetGauge(incCmd.Name)

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
