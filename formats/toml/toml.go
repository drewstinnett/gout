package toml

import (
	"context"

	utoml "github.com/pelletier/go-toml/v2"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return utoml.Marshal(v)
}

func (w Formatter) FormatWithContext(ctx context.Context, v interface{}) ([]byte, error) {
	return w.Format(v)
}
