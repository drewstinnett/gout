package gout

import (
	"github.com/drewstinnett/go-output-format/v2/formats/gotemplate"
	"github.com/drewstinnett/go-output-format/v2/formats/json"
	"github.com/drewstinnett/go-output-format/v2/formats/plain"
	"github.com/drewstinnett/go-output-format/v2/formats/yaml"
)

// BuiltInFormatters is a map of all formatters that we ship
var BuiltInFormatters = map[string]Formatter{
	"json":       json.Formatter{},
	"yaml":       yaml.Formatter{},
	"plain":      plain.Formatter{},
	"gotemplate": gotemplate.Formatter{},
}
