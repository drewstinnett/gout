package xml

import (
	"context"
	uxml "encoding/xml"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uxml.Marshal(v)
}

func (w Formatter) FormatWithContext(ctx context.Context, v interface{}) ([]byte, error) {
	return w.Format(v)
}
