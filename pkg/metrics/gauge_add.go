package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type GaugeAddHandler struct {
	registry *Registry
}

func NewGaugeAddHandler(registry *Registry) *GaugeAddHandler {
	return &GaugeAddHandler{registry: registry}
}

func (h *GaugeAddHandler) Handle(value interface{}) error {
	addCmd, ok := value.(*payload.GaugeAddCmd)

	if !ok {
		return errors.New("unexpected value for gauge add handler")
	}

	gauge, err := h.registry.GetGauge(addCmd.Name)

	if err != nil {
		return err
	}

	m, err := gauge.GetMetricWith(addCmd.Labels)

	if err != nil {
		return err
	}

	m.Add(addCmd.Value)

	return nil
}
