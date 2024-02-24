/*
Package yaml provides the yaml plugin for Gout
*/
package yaml

import (
	"github.com/drewstinnett/gout/v2/formats"
	uyaml "gopkg.in/yaml.v3"
)

// Formatter is the base structure that holds the yaml plugin stuff
type Formatter struct{}

// Format satisfies the formats.Formatter interface
func (w Formatter) Format(v any) ([]byte, error) {
	return uyaml.Marshal(v)
}

func init() {
	formats.Add("yaml", func() formats.Formatter {
		return &Formatter{}
	})
}
