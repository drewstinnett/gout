package json

import (
	ujson "encoding/json"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return ujson.Marshal(v)
}
