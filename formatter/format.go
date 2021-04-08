package formatter

import (
	"errors"
	"fmt"
)

// Starting with code similar to:
// https://github.com/hashicorp/vault/blob/master/command/format.go

type Formatter interface {
	Output(data interface{}) ([]byte, error)
	Format(data interface{}) ([]byte, error)
}

var Formatters = map[string]Formatter{
	"yaml": YamlFormatter{},
	"json": JsonFormatter{},
	//"plain": PlainFormatter{},
}

func GetFormats() []string {
	keys := make([]string, len(Formatters))

	i := 0
	for k := range Formatters {
		keys[i] = k
		i++
	}
	return keys
}

type FormatterConfig struct {
	Format string
}

func OutputData(data interface{}, config *FormatterConfig) ([]byte, error) {

	formatter, ok := Formatters[config.Format]
	if !ok {
		err := errors.New(fmt.Sprintf("Invalid output format: %s", config.Format))
		return nil, err
	}

	parsed, err := formatter.Output(data)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
