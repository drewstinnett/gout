package formats

import (
	"github.com/apex/log"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"gopkg.in/yaml.v2"
)

// YamlFormatter Basic YAML formatter struc
type YAMLFormatter struct{}

// Format How do we actually format YAML?
func (y YAMLFormatter) Format(data interface{}, config *config.Config) ([]byte, error) {
	log.Debugf("Called with config: %v", config)
	return yaml.Marshal(data)
}

// Output Do the output return string here
func (y YAMLFormatter) Output(data interface{}, config *config.Config) ([]byte, error) {
	log.Debugf("Called with config: %v", config)
	b, _ := y.Format(data, config)
	return b, nil
}
