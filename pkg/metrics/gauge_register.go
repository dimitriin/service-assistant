package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type GaugeRegisterHandler struct {
	registry *Registry
}

func NewGaugeRegisterHandler(registry *Registry) *GaugeRegisterHandler {
	return &GaugeRegisterHandler{registry: registry}
}

func (h *GaugeRegisterHandler) Handle(value interface{}) error {
	packetRegisterCmd, ok := value.(*payload.Packet_GaugeRegisterCmd)

	if !ok {
		return errors.New("unexpected value for gauge register handler")
	}

	registerCmd := packetRegisterCmd.GaugeRegisterCmd

	if err := h.registry.RegisterGauge(registerCmd.Name, registerCmd.Help, registerCmd.Labels); err != nil {
		return err
	}

	return nil
}
