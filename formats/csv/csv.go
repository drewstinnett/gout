package csv

import (
	"errors"

	"github.com/drewstinnett/go-output-format/v2/config"
	"github.com/jszwec/csvutil"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return csvutil.Marshal(v)
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}
