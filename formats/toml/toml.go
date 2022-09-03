package toml

import utoml "github.com/pelletier/go-toml/v2"

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return utoml.Marshal(v)
}
