package formatter

import (
	"gopkg.in/yaml.v2"
)

type YamlFormatter struct{}

func (y YamlFormatter) Format(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}

func (y YamlFormatter) Output(data interface{}) ([]byte, error) {
	b, err := y.Format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
