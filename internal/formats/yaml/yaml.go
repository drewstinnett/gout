package yaml

import (
	"github.com/apex/log"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"gopkg.in/yaml.v2"
)

// YamlFormatter Basic YAML formatter struc
type YAML struct{}

// Format How do we actually format YAML?
func (y *YAML) Format(data interface{}, config *config.Config) ([]byte, error) {
	log.Debugf("Called with config: %v", config)
	return yaml.Marshal(data)
}

// Output Do the output return string here
func (y *YAML) Output(data interface{}, config *config.Config) ([]byte, error) {
	log.Debugf("Called with config: %v", config)
	b, _ := y.Format(data, config)
	return b, nil
}

func init() {
	formatter.Add("yaml", func() formatter.Formatter {
		return &YAML{}
	})
}
