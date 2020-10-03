package config

type Histogram struct {
	Name    string `validate:"gt=0"`
	Help    string `validate:"gt=0"`
	Labels  []string
	Buckets []float64
}
