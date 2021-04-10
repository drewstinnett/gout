package formatter

import (
	"fmt"
)

// PlainFormatter Just output in raw go format
type plainFormatter struct{}

// Format Do the actual formatting here
func (p plainFormatter) format(data interface{}) ([]byte, error) {
	b := []byte(fmt.Sprintf("%+v", data.(interface{})))
	return b, nil
}

// Output Capture the output
func (p plainFormatter) output(data interface{}) ([]byte, error) {
	b, err := p.format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
