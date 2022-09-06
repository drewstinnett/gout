package plain

import (
	"context"
	"fmt"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v\n", v)), nil
}

func (w Formatter) FormatWithContext(ctx context.Context, v interface{}) ([]byte, error) {
	return w.Format(v)
}
