package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type GaugeSetHandler struct {
	registry *Registry
}

func NewGaugeSetHandler(registry *Registry) *GaugeSetHandler {
	return &GaugeSetHandler{registry: registry}
}

func (h *GaugeSetHandler) Handle(value interface{}) error {
	packetSetCmd, ok := value.(*payload.Packet_GaugeSetCmd)

	if !ok {
		return errors.New("unexpected value for gauge set handler")
	}

	setCmd := packetSetCmd.GaugeSetCmd

	gauge, err := h.registry.GetGauge(setCmd.Name)

	if err != nil {
		return err
	}

	m, err := gauge.GetMetricWith(setCmd.Labels)

	if err != nil {
		return err
	}

	m.Set(setCmd.Value)

	return nil
}
