package json

import (
	"testing"

	"github.com/drewstinnett/go-output-format/v2/config"
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
	o := config.FormatterOpts{
		"indent": "  ",
		"prefix": "+",
	}
	got, err := f.FormatWithOpts(struct {
		Foo string
	}{
		Foo: "bar",
	}, o)
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("{\n+  \"Foo\": \"bar\"\n+}"), string(got))
}
