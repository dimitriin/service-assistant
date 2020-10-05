package protocol

import (
	"fmt"
	"reflect"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type HandlerInterface interface {
	Handle(value interface{}) error
}

type PacketHandler struct {
	handlers map[string]HandlerInterface
}

func NewPacketHandler(handlers map[string]HandlerInterface) *PacketHandler {
	return &PacketHandler{handlers: handlers}
}

func (h *PacketHandler) Handle(packet *payload.Packet) error {
	handler, ok := h.handlers[reflect.TypeOf(packet.Payload).String()]

	if !ok {
		return fmt.Errorf("unexpected packet type %s", reflect.TypeOf(packet.Payload).String())
	}

	return handler.Handle(packet.Payload)
}
