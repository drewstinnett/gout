package gout

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestNewWithCobraCmdFlag(t *testing.T) {
	cmd := cobra.Command{}
	cmd.Flags().String("format", "yaml", "The format")
	c, err := NewWithCobraCmd(&cmd, nil)
	require.NoError(t, err)
	require.NotNil(t, c)
}

func TestNewWithCobraCmdPersistentFlag(t *testing.T) {
	cmd := cobra.Command{}
	cmd.PersistentFlags().String("format", "yaml", "The format")
	c, err := NewWithCobraCmd(&cmd, nil)
	require.NoError(t, err)
	require.NotNil(t, c)
}

func TestNewWithCobraCmdBadFormat(t *testing.T) {
	cmd := cobra.Command{}
	cmd.Flags().String("format", "not-exist", "The format")
	c, err := NewWithCobraCmd(&cmd, nil)
	require.EqualError(t, err, "Could not find the format not-exist")
	require.Nil(t, c)
}

func TestNewWithCobraCmdMissingFormat(t *testing.T) {
	cmd := cobra.Command{}
	c, err := NewWithCobraCmd(&cmd, nil)
	require.EqualError(t, err, "The flag 'format' is not available")
	require.Nil(t, c)
}
