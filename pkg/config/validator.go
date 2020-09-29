package config

import "gopkg.in/go-playground/validator.v9"

type Validator struct {}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(cfg Config) error {
	return validator.New().Struct(cfg)
}