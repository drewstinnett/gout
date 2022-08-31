package json

import ujson "encoding/json"

type JSONFormatter struct{}

func (w JSONFormatter) Format(v interface{}) ([]byte, error) {
	return ujson.Marshal(v)
}
