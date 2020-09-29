package config

type Counter struct {
	Name string `validate:"gt=0"`
	Namespace string
	Subsystem string
	Help string `validate:"gt=0"`
	Labels []string
}
