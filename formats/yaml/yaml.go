package yaml

import (
	"errors"

	"github.com/drewstinnett/go-output-format/v2/config"
	uyaml "gopkg.in/yaml.v3"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uyaml.Marshal(v)
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}
