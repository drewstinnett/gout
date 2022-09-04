package yaml

import (
	"testing"

	"github.com/drewstinnett/go-output-format/v2/config"
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

func TestYAMLFormatterWithOpts(t *testing.T) {
	f := Formatter{}
	opts := config.FormatterOpts{
		"indent": 10, // TODO: This does not work yet...y??
		"header": true,
	}
	got, err := f.FormatWithOpts([]struct {
		Foo string
	}{
		{Foo: "bar"},
		{Foo: "baz"},
	}, opts)
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("---\n- foo: bar\n- foo: baz\n"), string(got))
}
