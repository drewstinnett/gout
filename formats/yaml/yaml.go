package yaml

import (
	"context"

	uyaml "gopkg.in/yaml.v3"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uyaml.Marshal(v)
}

func (w Formatter) FormatWithContext(ctx context.Context, v interface{}) ([]byte, error) {
	return w.Format(v)
}
