package formatter

import (
	"fmt"
)

type PlainFormatter struct{}

func (j PlainFormatter) Format(data interface{}) ([]byte, error) {
	b := []byte(fmt.Sprintf("%+v", data.(interface{})))
	return b, nil
}

func (j PlainFormatter) Output(data interface{}) error {
	b, err := j.Format(data)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(b))
	return nil
}
