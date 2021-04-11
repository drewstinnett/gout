package formatter

import (
	"encoding/json"
)

// JSONFormatter Basic struct.
type jsonFormatter struct{}

// Format Do the formatting here.
func (j jsonFormatter) format(data interface{}, config *Config) ([]byte, error) {
	d, err := json.MarshalIndent(data, "", "  ")

	return d, err
}

// Output Capture output of JSON format
func (j jsonFormatter) output(data interface{}, config *Config) ([]byte, error) {
	b, err := j.format(data, config)
	if err != nil {
		return nil, err
	}
	return b, nil
}
