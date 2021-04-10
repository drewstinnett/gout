package formatter

import (
	"encoding/json"
)

// JSONFormatter Basic struct.
type JSONFormatter struct{}

// Format Do the formatting here.
func (j JSONFormatter) format(data interface{}) ([]byte, error) {
	d, err := json.MarshalIndent(data, "", "  ")

	return d, err
}

// Output Capture output of JSON format
func (j JSONFormatter) output(data interface{}) ([]byte, error) {
	b, err := j.format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
