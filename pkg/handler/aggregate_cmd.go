package handler

import (
	"fmt"

	"github.com/dimitriin/service-assistant/pkg/protocol/models"
)

type AggregateCmdHandler struct {
	counterIncHandler *CounterIncHandler
	counterAddHandler *CounterAddHandler
}

func NewAggregateCmdHandler(counterIncHandler *CounterIncHandler, counterAddHandler *CounterAddHandler) *AggregateCmdHandler {
	return &AggregateCmdHandler{counterIncHandler: counterIncHandler, counterAddHandler: counterAddHandler}
}

func (h *AggregateCmdHandler) Handle(value interface{}) error {
	switch cmd := value.(type) {
	case *models.CounterIncCmd:
		return h.counterIncHandler.handle(cmd)
	case *models.CounterAddCmd:
		return h.counterAddHandler.handle(cmd)
	default:
		return fmt.Errorf("unexpected cmd %#v", cmd)
	}
}
