package config

type Gauge struct {
	Name   string `validate:"gt=0"`
	Help   string `validate:"gt=0"`
	Labels []string
}
