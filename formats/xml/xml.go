package xml

import (
	uxml "encoding/xml"
	"errors"

	"github.com/drewstinnett/go-output-format/v2/config"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uxml.Marshal(v)
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}
