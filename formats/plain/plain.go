package plain

import (
	"fmt"

	"github.com/drewstinnett/gout/v2/formats"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v\n", v)), nil
}

func init() {
	formats.Add("plain", func() formats.Formatter {
		return &Formatter{}
	})
}
