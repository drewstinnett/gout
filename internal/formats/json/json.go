package json

import (
	"encoding/json"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

// JSONFormatter Basic struct.
type plug struct{}

// Format Do the formatting here.
func (p *plug) Format(data interface{}, config *config.Config) ([]byte, error) {
	d, err := json.MarshalIndent(data, "", "  ")

	return d, err
}

// Output Capture output of JSON format
func (p *plug) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, err := p.Format(data, config)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func init() {
	formatter.Add("json", func() formatter.Formatter {
		return &plug{}
	})
}
