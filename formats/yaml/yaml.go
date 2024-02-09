package yaml

import (
	uyaml "gopkg.in/yaml.v3"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uyaml.Marshal(v)
}
