package csv

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCSVFormatter(t *testing.T) {
	f := Formatter{}
	got, err := f.Format([]struct {
		Foo string
	}{
		{Foo: "bar"},
		{Foo: "baz"},
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("Foo\nbar\nbaz\n"), string(got))
}

func TestCSVFormatterWithContext(t *testing.T) {
	f := Formatter{}
	got, err := f.FormatWithContext(context.Background(), []struct {
		Foo string
	}{
		{Foo: "bar"},
		{Foo: "baz"},
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("Foo\nbar\nbaz\n"), string(got))
}
