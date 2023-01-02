package gout

import (
	"fmt"
	"strings"

	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats/gotemplate"
	"github.com/spf13/cobra"
)

// CobraCmdConfig defines what fields the formatting values are stored in
type CobraConfig struct {
	FormatField     string
	FormatDefault   string
	FormatHelp      string
	TemplateDefault string
	TemplateHelp    string
}

type CobraOption func(*CobraConfig)

func WithCobraFormatField(s string) CobraOption {
	return func(c *CobraConfig) {
		c.FormatField = s
	}
}

func WithCobraFormatDefault(s string) CobraOption {
	return func(c *CobraConfig) {
		c.FormatDefault = s
	}
}

func WithCobraTemplateDefault(s string) CobraOption {
	return func(c *CobraConfig) {
		c.TemplateDefault = s
	}
}

// WithCobraFormatHelp defines wht the help for the format itself should be
func WithCobraFormatHelp(s string) CobraOption {
	return func(c *CobraConfig) {
		c.FormatHelp = s
	}
}

// WithCobraTemplateHelp defines what the help for the format template should be
func WithCobraTemplateHelp(s string) CobraOption {
	return func(c *CobraConfig) {
		c.TemplateHelp = s
	}
}

// NewcobraCmdConfig generates a new CobraCmdConfig with some sane defaults
func NewCobraConfig(opts ...CobraOption) *CobraConfig {
	// Set up a default config
	cc := &CobraConfig{
		FormatField:     "format",
		FormatDefault:   "yaml",
		FormatHelp:      "Format to use for output",
		TemplateDefault: "{{ . }}",
		TemplateHelp:    "Template to use when using the gotemplate format",
	}

	// Override that stuff
	for _, opt := range opts {
		opt(cc)
	}

	return cc
}

// BindCobraCmd creates a new set of flags for Cobra that can be used to
// configure Gout
func BindCobra(cmd *cobra.Command, config *CobraConfig) error {
	if config == nil {
		config = NewCobraConfig()
	}
	keys := make([]string, 0, len(BuiltInFormatters))
	for k := range BuiltInFormatters {
		keys = append(keys, k)
	}
	help := config.FormatHelp + " (" + strings.Join(keys, "|") + ")"
	cmd.PersistentFlags().String(config.FormatField, config.FormatDefault, help)
	cmd.PersistentFlags().String(config.FormatField+"-template", config.TemplateDefault, config.TemplateHelp)
	return nil
}

// WithCobra sets up the the built-in Gout client using options from the cobra.Cmd
func WithCobra(cmd *cobra.Command, conf *CobraConfig) error {
	g := GetGout()
	err := ApplyCobra(g, cmd, conf)
	if err != nil {
		return err
	}
	return nil
}

// ApplyCobra uses settings from cobra.Command to set up a given Gout instance
func ApplyCobra(g *Gout, cmd *cobra.Command, conf *CobraConfig) error {
	if conf == nil {
		conf = NewCobraConfig()
	}

	var format string
	var err error
	format, err = cmd.Flags().GetString(conf.FormatField)
	if err != nil {
		return err
	}
	if format == "gotemplate" {
		t, err := cmd.PersistentFlags().GetString(conf.FormatField + "-template")
		if err != nil {
			return err
		}
		g.SetFormatter(gotemplate.Formatter{
			Opts: config.FormatterOpts{
				"template": t,
			},
		})
	} else {
		if fr, ok := BuiltInFormatters[format]; ok {
			g.SetFormatter(fr)
		} else {
			return fmt.Errorf("Could not find the format %v", format)
		}
	}
	return nil
}

// NewWithCobraCmd creates a pointer to a new writer with options from a cobra.Command
func NewWithCobraCmd(cmd *cobra.Command, conf *CobraConfig) (*Gout, error) {
	// Default this writer to stdout
	c, err := New()
	if err != nil {
		return nil, err
	}

	err = ApplyCobra(c, cmd, conf)
	if err != nil {
		return nil, err
	}

	return c, nil
}
