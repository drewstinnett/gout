/*
Package json provides a json output plugin for Gout
*/
package json

import (
	ujson "encoding/json"

	"github.com/drewstinnett/gout/v2/formats"
)

// Formatter holds the base json stuff
type Formatter struct {
	indent bool
}

// Format satisfies the formats.Formatter interface
func (w Formatter) Format(v any) ([]byte, error) {
	if w.indent {
		return ujson.MarshalIndent(v, "", "  ")
	}
	return ujson.Marshal(v)
}

func init() {
	formats.Add("json", func() formats.Formatter {
		return &Formatter{}
	})
}
