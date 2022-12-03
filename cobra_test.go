package gout

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

/*
func TestNewWithCobraCmdFlag(t *testing.T) {
	cmd := cobra.Command{}
	cmd.Flags().String("format", "yaml", "The format")
	c, err := NewWithCobraCmd(&cmd, nil)
	require.NoError(t, err)
	require.NotNil(t, c)
}
*/

func TestNewWithCobraCmdPersistentFlag(t *testing.T) {
	cmd := cobra.Command{}
	cmd.PersistentFlags().String("format", "yaml", "The format")
	err := cmd.Execute()
	require.NoError(t, err)
	c, err := NewWithCobraCmd(&cmd, nil)
	require.NoError(t, err)
	require.NotNil(t, c)
}

func TestNewWithCobraCmdPersistentFlagTemplate(t *testing.T) {
	cmd := cobra.Command{}
	cmd.PersistentFlags().String("format", "gotemplate", "The format")
	cmd.PersistentFlags().String("format-template", "{{ . }}", "The format template")
	err := cmd.Execute()
	require.NoError(t, err)
	c, err := NewWithCobraCmd(&cmd, nil)
	require.NoError(t, err)
	require.NotNil(t, c)
}

func TestNewWithCobraCmdBadFormat(t *testing.T) {
	cmd := cobra.Command{}
	cmd.PersistentFlags().String("format", "not-exist", "The format")
	err := cmd.Execute()
	require.NoError(t, err)
	c, err := NewWithCobraCmd(&cmd, nil)
	require.EqualError(t, err, "Could not find the format not-exist")
	require.Nil(t, c)
}

func TestNewWithCobraCmdMissingFormat(t *testing.T) {
	cmd := cobra.Command{}
	c, err := NewWithCobraCmd(&cmd, nil)
	require.EqualError(t, err, "flag accessed but not defined: format")
	require.Nil(t, c)
}

func TestBindCobraCmd(t *testing.T) {
	cmd := cobra.Command{}

	err := BindCobraCmd(&cmd, nil)
	require.NoError(t, err)

	got, err := cmd.PersistentFlags().GetString("format")
	require.NoError(t, err)
	require.Equal(t, "yaml", got)
}
