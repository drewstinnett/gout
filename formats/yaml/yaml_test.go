package yaml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Opts struct {
	items map[string]interface{}
}

func (o *Opts) Set(s string, v interface{}) error {
	o.items[s] = v
	return nil
}

func (o *Opts) Get(s string) (interface{}, error) {
	if _, ok := o.items[s]; ok {
		return nil, fmt.Errorf("Could not find config item '%s'", s)
	}
	return o.items[s], nil
}

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

/*
func TestYAMLFormatterWithOpts(t *testing.T) {
	f := Formatter{}
	got, err := f.FormatWithContext(context.Background(), []struct {
		Foo string
	}{
		{Foo: "bar"},
		{Foo: "baz"},
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("- foo: bar\n- foo: baz\n"), string(got))
}
*/

/*
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

*/
