package yaml

import (
	"bytes"
	"errors"

	"github.com/drewstinnett/go-output-format/v2/config"
	uyaml "gopkg.in/yaml.v3"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return uyaml.Marshal(v)
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	var b bytes.Buffer
	yamlEncoder := uyaml.NewEncoder(&b)
	if v, ok := o["header"]; ok {
		uh, ok := v.(bool)
		if !ok {
			return nil, errors.New(`"header" should be a bool`)
		}
		if uh {
			b.WriteString("---\n")
		}
	}
	if v, ok := o["indent"]; ok {
		i, ok := v.(int)
		if !ok {
			return nil, errors.New("Could not convert indent in to int")
		}
		yamlEncoder.SetIndent(i)
	}
	err := yamlEncoder.Encode(&v)
	return b.Bytes(), err
}
