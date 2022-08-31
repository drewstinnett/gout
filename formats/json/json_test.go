package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONFormatter(t *testing.T) {
	f := JSONFormatter{}
	got, err := f.Format(struct {
		Foo string
	}{
		Foo: "bar",
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string(`{"Foo":"bar"}`), string(got))
}
