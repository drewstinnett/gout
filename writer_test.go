package writer

import (
	"bytes"
	"testing"

	"github.com/drewstinnett/go-output-format/formats/plain"
	"github.com/stretchr/testify/require"
)

func TestNewWriter(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, c)
}

func TestWriterPrinter(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)
	var buf bytes.Buffer
	c.SetWriter(&buf)
	c.Print(struct{ Foo string }{Foo: "bar"})
	require.Equal(t, "foo: bar\n", buf.String())

	// Make sure we can change the formatter
	c.SetFormatter(plain.Formatter{})
	require.IsType(t, plain.Formatter{}, c.Formatter)
}

func TestPrintError(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)
	err = c.Print(make(chan int))
	require.Error(t, err)

	unprintable := make(chan int)
	require.Panics(t, func() { c.MustPrint(unprintable) })
}
