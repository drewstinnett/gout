package formatter

import (
	"fmt"
)

// PlainFormatter Just output in raw go format
type PlainFormatter struct{}

// Format Do the actual formatting here
func (j PlainFormatter) Format(data interface{}) ([]byte, error) {
	b := []byte(fmt.Sprintf("%+v", data.(interface{})))
	return b, nil
}

// Output Capture the output
func (j PlainFormatter) Output(data interface{}) ([]byte, error) {
	b, err := j.Format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
