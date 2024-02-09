package xml

import (
	uxml "encoding/xml"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uxml.Marshal(v)
}
