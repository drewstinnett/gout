package utils_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestSliceContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := utils.StringInSlice("bar", s)
	require.True(t, got)
}

func TestSliceNotContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := utils.StringInSlice("NeverExists", s)
	require.False(t, got)
}
