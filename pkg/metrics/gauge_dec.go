package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type GaugeDecHandler struct {
	registry *Registry
}

func NewGaugeDecHandler(registry *Registry) *GaugeDecHandler {
	return &GaugeDecHandler{registry: registry}
}

func (h *GaugeDecHandler) Handle(value interface{}) error {
	packetDecCmd, ok := value.(*payload.Packet_GaugeDecCmd)

	if !ok {
		return errors.New("unexpected value for gauge dec handler")
	}

	decCmd := packetDecCmd.GaugeDecCmd

	counter, err := h.registry.GetGauge(decCmd.Name)

	if err != nil {
		return err
	}

	m, err := counter.GetMetricWith(decCmd.Labels)

	if err != nil {
		return err
	}

	m.Dec()

	return nil
}
