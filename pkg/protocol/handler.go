package protocol

import (
	"fmt"
)

type HandlerInterface interface {
	Handle(value interface{}) error
}

type PacketHandler struct {
	handlers map[uint16]HandlerInterface
}

func NewPacketHandler(handlers map[uint16]HandlerInterface) *PacketHandler {
	return &PacketHandler{handlers: handlers}
}

func (h *PacketHandler) Handle(packet *Packet) error {
	handler, ok := h.handlers[packet.Type]

	if !ok {
		return fmt.Errorf("unexpected packet type %d", packet.Type)
	}

	return handler.Handle(packet.Value)
}
