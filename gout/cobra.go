package gout

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// CobraCmdConfig defines what fields the formatting values are stored in
type CobraCmdConfig struct {
	FormatField string
}

// NewClientWithCobraCmd creates a pointer to a new writer with options from a cobra.Command
func NewClientWithCobraCmd(cmd *cobra.Command, config *CobraCmdConfig) (*Client, error) {
	var err error
	if config == nil {
		config = &CobraCmdConfig{
			FormatField: "format",
		}
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
	c := &Client{}
	if fr, ok := BuiltInFormatters[format]; ok {
		c.SetFormatter(fr)
	} else {
		return nil, fmt.Errorf("Could not find the format %v", format)
	}
	return c, nil
}
