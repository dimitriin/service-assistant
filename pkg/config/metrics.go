package config

type Metrics struct {
	Counters   []Counter
	Histograms []Histogram
	Gauges     []Gauge
}
