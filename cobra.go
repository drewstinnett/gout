package gout

import (
	"fmt"
	"os"
	"strings"

	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats/gotemplate"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// CobraCmdConfig defines what fields the formatting values are stored in
type CobraCmdConfig struct {
	FormatField     string
	Default         string
	DefaultTemplate string
	Help            string
	HelpTemplate    string
}

// NewcobraCmdConfig generates a new CobraCmdConfig with some sane defaults
func NewCobraCmdConfig() *CobraCmdConfig {
	return &CobraCmdConfig{
		FormatField:     "format",
		Default:         "yaml",
		DefaultTemplate: "{{ . }}",
		Help:            "Format to use for output",
		HelpTemplate:    "Template to use when using the gotemplate format",
	}
}

// BindCobraCmd creates a new flag called 'format' that accepts the format.
func BindCobraCmd(cmd *cobra.Command, config *CobraCmdConfig) error {
	if config == nil {
		config = NewCobraCmdConfig()
	}
	keys := make([]string, 0, len(BuiltInFormatters))
	for k := range BuiltInFormatters {
		keys = append(keys, k)
	}
	help := config.Help + " (" + strings.Join(keys, "|") + ")"
	cmd.PersistentFlags().String(config.FormatField, config.Default, help)
	cmd.PersistentFlags().String(config.FormatField+"-template", config.DefaultTemplate, config.HelpTemplate)
	return nil
}

// NewWithCobraCmd creates a pointer to a new writer with options from a cobra.Command
func NewWithCobraCmd(cmd *cobra.Command, conf *CobraCmdConfig) (*Gout, error) {
	var err error
	if conf == nil {
		conf = NewCobraCmdConfig()
	}
	var format string

	format, err = cmd.Flags().GetString(conf.FormatField)
	if err != nil {
		return nil, err
	}

	// Default this writer to stdout
	c := &Gout{
		Writer: os.Stdout,
	}

	if format == "gotemplate" {
		t, err := cmd.PersistentFlags().GetString(conf.FormatField + "-template")
		if err != nil {
			log.Warn().Err(err).Msg("Error looking up template, using the default instead")
		}
		c.SetFormatter(gotemplate.Formatter{
			Opts: config.FormatterOpts{
				"template": t,
			},
			// Template: t,
		})
	} else {
		if fr, ok := BuiltInFormatters[format]; ok {
			c.SetFormatter(fr)
		} else {
			return nil, fmt.Errorf("Could not find the format %v", format)
		}
	}
	return c, nil
}
