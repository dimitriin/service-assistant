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
	packetSubCmd, ok := value.(*payload.Packet_GaugeSubCmd)

	if !ok {
		return errors.New("unexpected value for gauge sub handler")
	}

	subCmd := packetSubCmd.GaugeSubCmd

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
