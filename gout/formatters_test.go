package gout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuiltInFormatters(t *testing.T) {
	require.Contains(t, BuiltInFormatters, "yaml")
	require.NotContains(t, BuiltInFormatters, "never-exist")
}
