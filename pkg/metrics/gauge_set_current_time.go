package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type GaugeSetToCurrentTimeHandler struct {
	registry *Registry
}

func NewGaugeSetToCurrentTimeHandler(registry *Registry) *GaugeSetToCurrentTimeHandler {
	return &GaugeSetToCurrentTimeHandler{registry: registry}
}

func (h *GaugeSetToCurrentTimeHandler) Handle(value interface{}) error {
	packetSetCmd, ok := value.(*payload.Packet_GaugeSetToCurrentTimeCmd)

	if !ok {
		return errors.New("unexpected value for gauge set to current time handler")
	}

	setCmd := packetSetCmd.GaugeSetToCurrentTimeCmd

	gauge, err := h.registry.GetGauge(setCmd.Name)

	if err != nil {
		return err
	}

	m, err := gauge.GetMetricWith(setCmd.Labels)

	if err != nil {
		return err
	}

	m.SetToCurrentTime()

	return nil
}
