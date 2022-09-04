package toml

import (
	"errors"

	"github.com/drewstinnett/go-output-format/v2/config"
	utoml "github.com/pelletier/go-toml/v2"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return utoml.Marshal(v)
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}
