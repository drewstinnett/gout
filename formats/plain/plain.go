package plain

import "fmt"

type PlainFormatter struct{}

func (w PlainFormatter) Format(v interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v\n", v)), nil
}
