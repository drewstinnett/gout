package formatter

import (
	"fmt"
)

// Starting with code similar to:
// https://github.com/hashicorp/vault/blob/master/command/format.go

// Formatter Generic formatter interface, all interfaces should provide a
// structure like this. All formatters must implement an output and format
// function. Output will do the actual data output, running Format first, to do
// the actual data formatting
type formatter interface {
	output(data interface{}, config *Config) ([]byte, error)
	format(data interface{}, config *Config) ([]byte, error)
}

// formatters Map of the different types of formatting we do here. The
// formatter must be registered in this map to be available
var formatters = map[string]formatter{
	"yaml":       yamlFormatter{},
	"json":       jsonFormatter{},
	"tsv":        tsvFormatter{},
	"plain":      plainFormatter{},
	"gotemplate": gotemplateFormatter{},
}

// GetFormats Return a list of formats available in formatters. Useful if you
// need to check what formatters are available in a standardized way
func GetFormats() []string {
	keys := make([]string, len(formatters))

	i := 0
	for k := range formatters {
		keys[i] = k
		i++
	}
	return keys
}

// Config Structure to pass to formatters.  Should include enough config to do
// the output. You must set the Format here to something like yaml, json,
// plain, or any other value returned by the GetFormats function
type Config struct {
	Format      string
	LimitFields []string
	Template    string
}

// OutputData Main function to return the data we will be printing to the
// screen. This is where the magic happens!
func OutputData(data interface{}, config *Config) ([]byte, error) {

	formatter, ok := formatters[config.Format]
	if !ok {
		err := fmt.Errorf("Invalid output format: %s", config.Format)
		return nil, err
	}

	parsed, err := formatter.output(data, config)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
