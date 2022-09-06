package json

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONFormatter(t *testing.T) {
	f := Formatter{}
	got, err := f.Format(struct {
		Foo string
	}{
		Foo: "bar",
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string(`{"Foo":"bar"}`), string(got))
}

func TestJSONFormatterWithOpts(t *testing.T) {
	f := Formatter{}
	ctx := context.WithValue(context.Background(), IndentField{}, "yes")
	got, err := f.FormatWithContext(ctx, struct {
		Foo string
	}{
		Foo: "bar",
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("{\n  \"Foo\": \"bar\"\n}"), string(got))
}
