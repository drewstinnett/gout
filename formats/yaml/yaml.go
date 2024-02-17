package yaml

import (
	"github.com/drewstinnett/gout/v2/formats"
	uyaml "gopkg.in/yaml.v3"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uyaml.Marshal(v)
}

func init() {
	formats.Add("yaml", func() formats.Formatter {
		return &Formatter{}
	})
}
