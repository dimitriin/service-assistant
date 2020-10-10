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
	packetIncCmd, ok := value.(*payload.Packet_GaugeIncCmd)

	if !ok {
		return errors.New("unexpected value for gauge inc handler")
	}

	incCmd := packetIncCmd.GaugeIncCmd

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
