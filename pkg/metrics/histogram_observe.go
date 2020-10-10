package metrics

import (
	"errors"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
)

type HistogramObserveHandler struct {
	registry *Registry
}

func NewHistogramObserveHandler(registry *Registry) *HistogramObserveHandler {
	return &HistogramObserveHandler{registry: registry}
}

func (h *HistogramObserveHandler) Handle(value interface{}) error {
	packetObserveCmd, ok := value.(*payload.Packet_HistogramObserveCmd)

	if !ok {
		return errors.New("unexpected value for histogram observe handler")
	}

	observeCmd := packetObserveCmd.HistogramObserveCmd

	histogram, err := h.registry.GetHistogram(observeCmd.Name)

	if err != nil {
		return err
	}

	m, err := histogram.GetMetricWith(observeCmd.Labels)

	if err != nil {
		return err
	}

	m.Observe(observeCmd.Value)

	return nil
}
