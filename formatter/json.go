package formatter

import (
	"encoding/json"
)

// JSONFormatter Basic struct.
type jsonFormatter struct{}

// Format Do the formatting here.
func (j jsonFormatter) format(data interface{}) ([]byte, error) {
	d, err := json.MarshalIndent(data, "", "  ")

	return d, err
}

// Output Capture output of JSON format
func (j jsonFormatter) output(data interface{}) ([]byte, error) {
	b, err := j.format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
