package formatter

import (
	"encoding/json"
)

type JsonFormatter struct{}

func (j JsonFormatter) Format(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

func (j JsonFormatter) Output(data interface{}) ([]byte, error) {
	b, err := j.Format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
