package handler

import (
	"fmt"

	"github.com/dimitriin/service-assistant/pkg/protocol/cmd"
)

type AggregateCmdHandler struct {
	counterIncHandler *CounterIncHandler
	counterAddHandler *CounterAddHandler
}

func NewAggregateCmdHandler(counterIncHandler *CounterIncHandler, counterAddHandler *CounterAddHandler) *AggregateCmdHandler {
	return &AggregateCmdHandler{counterIncHandler: counterIncHandler, counterAddHandler: counterAddHandler}
}

func (h *AggregateCmdHandler) Handle(value interface{}) error {
	switch cmdValue := value.(type) {
	case *cmd.CounterIncCmd:
		return h.counterIncHandler.handle(cmdValue)
	case *cmd.CounterAddCmd:
		return h.counterAddHandler.handle(cmdValue)
	default:
		return fmt.Errorf("unexpected cmd %#v", cmdValue)
	}
}
