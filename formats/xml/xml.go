package xml

import (
	uxml "encoding/xml"

	"github.com/drewstinnett/gout/v2/formats"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uxml.Marshal(v)
}

func init() {
	formats.Add("xml", func() formats.Formatter {
		return &Formatter{}
	})
}
