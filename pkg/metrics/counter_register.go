package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type CounterRegisterHandler struct {
	registry *Registry
}

func NewCounterRegisterHandler(registry *Registry) *CounterRegisterHandler {
	return &CounterRegisterHandler{registry: registry}
}

func (h *CounterRegisterHandler) Handle(value interface{}) error {
	packetRegisterCmd, ok := value.(*payload.Packet_CounterRegisterCmd)

	if !ok {
		return errors.New("unexpected value for counter register handler")
	}

	registerCmd := packetRegisterCmd.CounterRegisterCmd

	if err := h.registry.RegisterCounter(registerCmd.Name, registerCmd.Help, registerCmd.Labels); err != nil {
		return err
	}

	return nil
}
