package yaml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYAMLFormatter(t *testing.T) {
	f := Formatter{}
	got, err := f.Format(struct {
		Foo string
	}{
		Foo: "bar",
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("foo: bar\n"), string(got))
}
