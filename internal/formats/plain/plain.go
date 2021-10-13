package plain

import (
	"fmt"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

// PlainFormatter Just output in raw go format
type Formatter struct{}

// Format Do the actual formatting here
func (p Formatter) Format(data interface{}, config *config.Config) ([]byte, error) {
	// b := []byte(fmt.Sprintf("%+v", data.(interface{})))
	b := []byte(fmt.Sprintf("%+v", data))
	return b, nil
}

// Output Capture the output
func (p Formatter) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, _ := p.Format(data, config)
	return b, nil
}

func init() {
	formatter.Add("plain", func() formatter.Formatter {
		return &Formatter{}
	})
}
