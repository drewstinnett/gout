package formatter

import (
	//_ "github.com/drewstinnett/go-output-format/internal/formats/yaml"
	"fmt"

	"github.com/drewstinnett/go-output-format/internal/utils"
	"github.com/drewstinnett/go-output-format/pkg/config"
)

// Starting with code similar to:
// https://github.com/hashicorp/vault/blob/master/command/format.go

// Formatter Generic formatter interface, all interfaces should provide a
// structure like this. All formatters must implement an output and format
// function. Output will do the actual data output, running Format first, to do
// the actual data formatting
type Formatter interface {
	Format(data interface{}, config *config.Config) ([]byte, error)
	Output(data interface{}, config *config.Config) ([]byte, error)
}

type Format interface{}

// GetFormats Return a list of formats available in formatters. Useful if you
// need to check what formatters are available in a standardized way
func GetFormats() []string {
	keys := make([]string, len(Formats))

	i := 0
	for k := range Formats {
		keys[i] = k
		i++
	}
	return keys
}

// OutputData Main function to return the data we will be printing to the
// screen. This is where the magic happens!
func OutputData(data interface{}, config *config.Config) ([]byte, error) {
	// Make sure it's a valid format
	if !utils.StringInSlice(config.Format, GetFormats()) {
		err := fmt.Errorf("Invalid output format: %s", config.Format)
		return nil, err
	}
	formatter := Formats[config.Format]()

	parsed, err := formatter.Output(data, config)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

type Creator func() Formatter

var Formats = map[string]Creator{}

func Add(name string, creator Creator) {
	Formats[name] = creator
}
