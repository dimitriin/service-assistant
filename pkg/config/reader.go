package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Reader struct {
	path string
}

func NewReader(path string) *Reader {
	return &Reader{path: path}
}

func (r *Reader) Read(cfg *Config) error {
	data, err := ioutil.ReadFile(r.path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		return err
	}

	return nil
}
