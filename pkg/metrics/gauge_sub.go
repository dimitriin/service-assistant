package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type GaugeSubHandler struct {
	registry *Registry
}

func NewGaugeSubHandler(registry *Registry) *GaugeSubHandler {
	return &GaugeSubHandler{registry: registry}
}

func (h *GaugeSubHandler) Handle(value interface{}) error {
	subCmd, ok := value.(*payload.GaugeSubCmd)

	if !ok {
		return errors.New("unexpected value for gauge sub handler")
	}

	gauge, err := h.registry.GetGauge(subCmd.Name)

	if err != nil {
		return err
	}

	m, err := gauge.GetMetricWith(subCmd.Labels)

	if err != nil {
		return err
	}

	m.Sub(subCmd.Value)

	return nil
}
