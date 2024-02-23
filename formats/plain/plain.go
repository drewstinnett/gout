/*
Package plain just prints the variable as %+v
*/
package plain

import (
	"fmt"

	"github.com/drewstinnett/gout/v2/formats"
)

// Formatter is the base struct for the plain plugin
type Formatter struct{}

// Format satisfies the formats.Formatter interface
func (w Formatter) Format(v interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v\n", v)), nil
}

func init() {
	formats.Add("plain", func() formats.Formatter {
		return &Formatter{}
	})
}
