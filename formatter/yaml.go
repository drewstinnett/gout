package formatter

import (
	"gopkg.in/yaml.v2"
)

// YamlFormatter Basic YAML formatter struc
type YamlFormatter struct{}

// Format How do we actually format YAML?
func (y YamlFormatter) Format(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}

// Output Do the output return string here
func (y YamlFormatter) Output(data interface{}) ([]byte, error) {
	b, err := y.Format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
