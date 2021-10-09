package formats

import (
	"fmt"

	"github.com/drewstinnett/go-output-format/pkg/config"
)

// PlainFormatter Just output in raw go format
type PlainFormatter struct{}

// Format Do the actual formatting here
func (p PlainFormatter) Format(data interface{}, config *config.Config) ([]byte, error) {
	// b := []byte(fmt.Sprintf("%+v", data.(interface{})))
	b := []byte(fmt.Sprintf("%+v", data))
	return b, nil
}

// Output Capture the output
func (p PlainFormatter) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, _ := p.Format(data, config)
	return b, nil
}
