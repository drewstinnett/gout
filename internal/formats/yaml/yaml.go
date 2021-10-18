package yaml

import (
	"github.com/apex/log"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"gopkg.in/yaml.v2"
)

// YamlFormatter Basic YAML formatter struc
type plug struct{}

// Format How do we actually format YAML?
func (p *plug) Format(data interface{}, config *config.Config) ([]byte, error) {
	log.Debugf("Called with config: %v", config)
	return yaml.Marshal(data)
}

// Output Do the output return string here
func (p *plug) Output(data interface{}, config *config.Config) ([]byte, error) {
	log.Debugf("Called with config: %v", config)
	b, _ := p.Format(data, config)
	return b, nil
}

func init() {
	formatter.Add("yaml", func() formatter.Formatter {
		return &plug{}
	})
}
