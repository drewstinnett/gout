package plain

import "fmt"

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v\n", v)), nil
}
