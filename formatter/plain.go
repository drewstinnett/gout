package formatter

import (
	"fmt"
)

// PlainFormatter Just output in raw go format
type PlainFormatter struct{}

// Format Do the actual formatting here
func (p PlainFormatter) format(data interface{}) ([]byte, error) {
	b := []byte(fmt.Sprintf("%+v", data.(interface{})))
	return b, nil
}

// Output Capture the output
func (p PlainFormatter) output(data interface{}) ([]byte, error) {
	b, err := p.format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
