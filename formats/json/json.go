/*
Package json provides a json output plugin for Gout
*/
package json

import (
	ujson "encoding/json"

	"github.com/drewstinnett/gout/v2/formats"
)

// Formatter holds the base json stuff
type Formatter struct{}

// Format satiesfiles the formats.Formatter interface
func (w Formatter) Format(v interface{}) ([]byte, error) {
	var i any
	if i == nil {
		return ujson.Marshal(v)
	}
	return ujson.MarshalIndent(v, "", "  ")
}

func init() {
	formats.Add("json", func() formats.Formatter {
		return &Formatter{}
	})
}
