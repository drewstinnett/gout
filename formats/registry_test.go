package formats_test

import (
	"testing"

	"github.com/drewstinnett/gout/v2/formats"
	_ "github.com/drewstinnett/gout/v2/formats/builtin"
	"github.com/stretchr/testify/require"
)

func TestBuiltinRegistry(t *testing.T) {
	require.Equal(
		t,
		[]string{"gotemplate", "json", "yaml"},
		formats.Names(),
	)
}
