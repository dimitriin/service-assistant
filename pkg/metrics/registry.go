package metrics

import (
	"errors"
	"sort"

	"github.com/dimitriin/service-assistant/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
)

var CounterNotRegisteredErr = errors.New("counter not registered")
var HistogramNotRegisteredErr = errors.New("histogram not registered")
var GaugeNotRegisteredErr = errors.New("gauge not registered")

type Registry struct {
	cfg        config.Metrics
	counters   map[string]*prometheus.CounterVec
	histograms map[string]*prometheus.HistogramVec
	gauges     map[string]*prometheus.GaugeVec
}

func NewRegistry(cfg config.Metrics) *Registry {
	return &Registry{
		cfg:        cfg,
		counters:   map[string]*prometheus.CounterVec{},
		histograms: map[string]*prometheus.HistogramVec{},
		gauges:     map[string]*prometheus.GaugeVec{},
	}
}

func (r *Registry) Register() error {
	for _, counterCfg := range r.cfg.Counters {
		if err := r.RegisterCounter(counterCfg.Name, counterCfg.Help, counterCfg.Labels); err != nil {
			return err
		}
	}

	for _, histogramCfg := range r.cfg.Histograms {
		if err := r.RegisterHistogram(histogramCfg.Name, histogramCfg.Help, histogramCfg.Labels, histogramCfg.Buckets); err != nil {
			return err
		}
	}

	for _, gaugeCfg := range r.cfg.Gauges {
		if err := r.RegisterGauge(gaugeCfg.Name, gaugeCfg.Help, gaugeCfg.Labels); err != nil {
			return err
		}
	}

	return nil
}

func (r *Registry) RegisterCounter(name, help string, labels []string) error {
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:      name,
		Namespace: "",
		Subsystem: "",
		Help:      help,
	}, labels)

	err := prometheus.Register(counter)

	if err != nil {
		return err
	}

	r.counters[name] = counter

	return nil
}

func (r *Registry) GetCounter(name string) (*prometheus.CounterVec, error) {
	counter, ok := r.counters[name]

	if !ok {
		return nil, CounterNotRegisteredErr
	}

	return counter, nil
}

func (r *Registry) RegisterHistogram(name, help string, labels []string, buckets []float64) error {
	sort.Float64s(buckets)

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:      name,
		Namespace: "",
		Subsystem: "",
		Help:      help,
		Buckets:   buckets,
	}, labels)

	err := prometheus.Register(histogram)

	if err != nil {
		return err
	}

	r.histograms[name] = histogram

	return nil
}

func (r *Registry) GetHistogram(name string) (*prometheus.HistogramVec, error) {
	histogram, ok := r.histograms[name]

	if !ok {
		return nil, HistogramNotRegisteredErr
	}

	return histogram, nil
}

func (r *Registry) RegisterGauge(name, help string, labels []string) error {
	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}, labels)

	if err := prometheus.Register(gauge); err != nil {
		return err
	}

	r.gauges[name] = gauge

	return nil
}

func (r *Registry) GetGauge(name string) (*prometheus.GaugeVec, error) {
	gauge, ok := r.gauges[name]

	if !ok {
		return nil, GaugeNotRegisteredErr
	}

	return gauge, nil
}
