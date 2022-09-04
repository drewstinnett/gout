package plain

import (
	"errors"
	"fmt"

	"github.com/drewstinnett/go-output-format/v2/config"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v\n", v)), nil
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}
