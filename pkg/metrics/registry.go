package metrics

import (
	"errors"
	"fmt"

	"github.com/dimitriin/service-assistant/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
)

var CounterNotRegisteredErr = errors.New("counter not registered")

type Registry struct {
	cfg      config.Metrics
	counters map[string]*prometheus.CounterVec
}

func NewRegistry(cfg config.Metrics) *Registry {
	return &Registry{cfg: cfg, counters: map[string]*prometheus.CounterVec{}}
}

func (r *Registry) Register() error {
	for _, counterCfg := range r.cfg.Counters {
		counter := prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:      counterCfg.Name,
			Namespace: counterCfg.Namespace,
			Subsystem: counterCfg.Subsystem,
			Help:      counterCfg.Help,
		}, counterCfg.Labels)

		err := prometheus.Register(counter)

		if err != nil {
			return err
		}

		r.counters[r.getCounterKey(counterCfg)] = counter
	}

	return nil
}

func (r *Registry) GetCounter(name string) (*prometheus.CounterVec, error) {
	counter, ok := r.counters[name]

	if !ok {
		return nil, CounterNotRegisteredErr
	}

	return counter, nil
}

func (r *Registry) getCounterKey(counterCfg config.Counter) string {
	key := counterCfg.Name

	if len(counterCfg.Subsystem) > 0 {
		key = fmt.Sprintf("%s_%s", counterCfg.Subsystem, key)
	}

	if len(counterCfg.Namespace) > 0 {
		key = fmt.Sprintf("%s_%s", counterCfg.Namespace, key)
	}

	return key
}
