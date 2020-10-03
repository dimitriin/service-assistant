package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type HistogramRegisterHandler struct {
	registry *Registry
}

func NewHistogramRegisterHandler(registry *Registry) *HistogramRegisterHandler {
	return &HistogramRegisterHandler{registry: registry}
}

func (h *HistogramRegisterHandler) Handle(value interface{}) error {
	registerCmd, ok := value.(*payload.HistogramRegisterCmd)

	if !ok {
		return errors.New("unexpected value for histogram register handler")
	}

	if err := h.registry.RegisterHistogram(registerCmd.Name, registerCmd.Help, registerCmd.Labels, registerCmd.Buckets); err != nil {
		return err
	}

	return nil
}
