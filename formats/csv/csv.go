package csv

import (
	"context"

	"github.com/jszwec/csvutil"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return csvutil.Marshal(v)
}

func (w Formatter) FormatWithContext(ctx context.Context, v interface{}) ([]byte, error) {
	return w.Format(v)
}
