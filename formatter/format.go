package formatter

import (
	"fmt"
)

// Starting with code similar to:
// https://github.com/hashicorp/vault/blob/master/command/format.go

// Formatter Generic formatter interface, all interfaces should provide a structure like this
type Formatter interface {
	Output(data interface{}) ([]byte, error)
	Format(data interface{}) ([]byte, error)
}

// Formatters Map of the different types of formatting we do here
var Formatters = map[string]Formatter{
	"yaml": YamlFormatter{},
	"json": JSONFormatter{},
	"tsv":  TsvFormatter{},
	//"plain": PlainFormatter{},
}

// GetFormats Return a list of formats available in Formatters
func GetFormats() []string {
	keys := make([]string, len(Formatters))

	i := 0
	for k := range Formatters {
		keys[i] = k
		i++
	}
	return keys
}

// Config Structure to pass to formatters.  Should include enough config to do the output
type Config struct {
	Format string
}

// OutputData Main function to return the data we will be printing to the screen
func OutputData(data interface{}, config *Config) ([]byte, error) {

	formatter, ok := Formatters[config.Format]
	if !ok {
		err := fmt.Errorf("Invalid output format: %s", config.Format)
		return nil, err
	}

	parsed, err := formatter.Output(data)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
