/*
Package xml provides an XML plugin for Gout
*/
package xml

import (
	uxml "encoding/xml"

	"github.com/drewstinnett/gout/v2/formats"
)

// Formatter is the base struct for the xml plugin
type Formatter struct{}

// Format satisfies the formats.Formatter interface
func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uxml.Marshal(v)
}

func init() {
	formats.Add("xml", func() formats.Formatter {
		return &Formatter{}
	})
}
