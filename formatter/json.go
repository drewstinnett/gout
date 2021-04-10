package formatter

import (
	"encoding/json"
)

// JSONFormatter Basic struct.
type JSONFormatter struct{}

// Format Do the formatting here.
func (j JSONFormatter) Format(data interface{}) ([]byte, error) {
	d, err := json.MarshalIndent(data, "", "  ")

	return d, err
}

// Output Capture output of JSON format
func (j JSONFormatter) Output(data interface{}) ([]byte, error) {
	b, err := j.Format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
