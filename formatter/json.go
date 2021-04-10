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
	b, _ := j.format(data, config)
	return b, nil
}
