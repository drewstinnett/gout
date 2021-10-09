package formats

import (
	"encoding/json"

	"github.com/drewstinnett/go-output-format/pkg/config"
)

// JSONFormatter Basic struct.
type JSONFormatter struct{}

// Format Do the formatting here.
func (j JSONFormatter) Format(data interface{}, config *config.Config) ([]byte, error) {
	d, err := json.MarshalIndent(data, "", "  ")

	return d, err
}

// Output Capture output of JSON format
func (j JSONFormatter) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, err := j.Format(data, config)
	if err != nil {
		return nil, err
	}
	return b, nil
}
