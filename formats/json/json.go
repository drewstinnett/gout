package json

import (
	ujson "encoding/json"

	"github.com/drewstinnett/go-output-format/v2/config"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return ujson.Marshal(v)
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	var prefix, indent string
	if v, ok := o["indent"]; ok {
		indent = v.(string)
	}
	if v, ok := o["prefix"]; ok {
		prefix = v.(string)
	}
	return ujson.MarshalIndent(v, prefix, indent)
}
