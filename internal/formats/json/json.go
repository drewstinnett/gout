package json

import (
	"encoding/json"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

// JSONFormatter Basic struct.
type Formatter struct{}

// Format Do the formatting here.
func (j *Formatter) Format(data interface{}, config *config.Config) ([]byte, error) {
	d, err := json.MarshalIndent(data, "", "  ")

	return d, err
}

// Output Capture output of JSON format
func (j *Formatter) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, err := j.Format(data, config)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func init() {
	formatter.Add("json", func() formatter.Formatter {
		return &Formatter{}
	})
}
