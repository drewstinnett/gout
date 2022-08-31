package tsv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTSVFormatter(t *testing.T) {
	f := TSVFormatter{}
	got, err := f.Format(struct {
		Foo string
		Baz string
	}{
		Foo: "bar",
		Baz: "bazinga",
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("bazinga\tbar\t\n"), string(got))
}
