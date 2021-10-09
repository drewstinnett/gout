package formatter

import (
	"fmt"

	"github.com/drewstinnett/go-output-format/internal/formats"
	"github.com/drewstinnett/go-output-format/pkg/config"
)

// Starting with code similar to:
// https://github.com/hashicorp/vault/blob/master/command/format.go

// Formatter Generic formatter interface, all interfaces should provide a
// structure like this. All formatters must implement an output and format
// function. Output will do the actual data output, running Format first, to do
// the actual data formatting
type formatter interface {
	Output(data interface{}, config *config.Config) ([]byte, error)
	Format(data interface{}, config *config.Config) ([]byte, error)
}

// formatters Map of the different types of formatting we do here. The
// formatter must be registered in this map to be available
var formatters = map[string]formatter{
	"yaml":       formats.YAMLFormatter{},
	"json":       formats.JSONFormatter{},
	"tsv":        formats.TSVFormatter{},
	"plain":      formats.PlainFormatter{},
	"gotemplate": formats.GoTemplateFormatter{},
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

// OutputData Main function to return the data we will be printing to the
// screen. This is where the magic happens!
func OutputData(data interface{}, config *config.Config) ([]byte, error) {
	formatter, ok := formatters[config.Format]
	if !ok {
		err := fmt.Errorf("Invalid output format: %s", config.Format)
		return nil, err
	}

	parsed, err := formatter.Output(data, config)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
