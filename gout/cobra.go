package gout

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// CobraCmdConfig defines what fields the formatting values are stored in
type CobraCmdConfig struct {
	FormatField string
	Default     string
	Help        string
}

// NewcobraCmdConfig generates a new CobraCmdConfig with some sane defaults
func NewCobraCmdConfig() *CobraCmdConfig {
	return &CobraCmdConfig{
		FormatField: "format",
		Default:     "yaml",
		Help:        "Format to use for output",
	}
}

// BindCobraCmd creates a new flag called 'format' that accepts the format.
func BindCobraCmd(cmd *cobra.Command, config *CobraCmdConfig) error {
	if config == nil {
		config = NewCobraCmdConfig()
	}
	cmd.PersistentFlags().String(config.FormatField, config.Default, config.Help)
	return nil
}

// NewClientWithCobraCmd creates a pointer to a new writer with options from a cobra.Command
func NewWithCobraCmd(cmd *cobra.Command, config *CobraCmdConfig) (*Client, error) {
	var err error
	if config == nil {
		config = NewCobraCmdConfig()
	}
	var format string

	// Check all the flags for a matching format
	for _, fs := range []pflag.FlagSet{
		*cmd.Flags(),
		*cmd.PersistentFlags(),
	} {
		format, err = fs.GetString(config.FormatField)
		if err == nil {
			break
		}
	}
	// If still not found, error
	if err != nil {
		return nil, fmt.Errorf("The flag '%v' is not available", config.FormatField)
	}
	// Default this writer to stdout
	c := &Client{
		Writer: os.Stdout,
	}
	if fr, ok := BuiltInFormatters[format]; ok {
		c.SetFormatter(fr)
	} else {
		return nil, fmt.Errorf("Could not find the format %v", format)
	}
	return c, nil
}
