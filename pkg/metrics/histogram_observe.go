package metrics

import "github.com/dimitriin/service-assistant/pkg/protocol/payload"

type HistogramObserveHandler struct {
	registry *Registry
}

func NewHistogramObserveHandler(registry *Registry) *HistogramObserveHandler {
	return &HistogramObserveHandler{registry: registry}
}

func (h *HistogramObserveHandler) Handle(value interface{}) error {
	observeCmd := value.(*payload.HistogramObserveCmd)

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
