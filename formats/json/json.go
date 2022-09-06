package json

import (
	"context"
	ujson "encoding/json"
)

type Formatter struct {
	ctx context.Context
}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	var i any
	if w.ctx != nil {
		i = w.ctx.Value("indent")
	}
	if i == nil {
		return ujson.Marshal(v)
	}
	return ujson.MarshalIndent(v, "", "  ")
}

func (w Formatter) FormatWithContext(ctx context.Context, v interface{}) ([]byte, error) {
	return w.withContext(ctx).Format(v)
}

func (w *Formatter) withContext(ctx context.Context) *Formatter {
	w.ctx = ctx
	return w
}
