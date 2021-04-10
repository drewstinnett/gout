package formatter

import (
	"gopkg.in/yaml.v2"
)

// YamlFormatter Basic YAML formatter struc
type yamlFormatter struct{}

// Format How do we actually format YAML?
func (y yamlFormatter) format(data interface{}, config *Config) ([]byte, error) {
	return yaml.Marshal(data)
}

// Output Do the output return string here
func (y yamlFormatter) output(data interface{}, config *Config) ([]byte, error) {
	b, err := y.format(data, config)
	if err != nil {
		return nil, err
	}
	return b, nil
}
