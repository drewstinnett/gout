package gout

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/drewstinnett/gout/v2/formats/json"
	"github.com/drewstinnett/gout/v2/formats/plain"
	"github.com/stretchr/testify/require"
)

func TestNewWriter(t *testing.T) {
	require.NotNil(t, New())
}

func TestSetGoutWriter(t *testing.T) {
	c := New()
	var buf bytes.Buffer
	c.SetWriter(&buf)
	c.Print(struct{ Foo string }{Foo: "bar"})
	require.Equal(t, "foo: bar\n", buf.String())

	// Make sure we can change the formatter
	c.SetFormatter(plain.Formatter{})
	require.IsType(t, plain.Formatter{}, c.formatter)
}

func TestNewGoutWithWriter(t *testing.T) {
	var b bytes.Buffer
	c := New(WithWriter(&b))
	require.NotNil(t, c)
	c.Print(struct{ Foo string }{Foo: "bar"})
	require.Equal(t, "foo: bar\n", b.String())
}

func TestNewGoutWithFormatter(t *testing.T) {
	var b bytes.Buffer
	c := New(WithWriter(&b), WithFormatter(plain.Formatter{}))
	require.NotNil(t, c)
	c.Print(struct{ Foo string }{Foo: "bar"})
	require.Equal(t, "{Foo:bar}\n", b.String())
}

func TestPrintError(t *testing.T) {
	c := New()
	require.Error(t, c.Print(make(chan int)))

	unprintable := make(chan int)
	require.Panics(t, func() { c.MustPrint(unprintable) })
	require.Panics(t, func() { c.MustPrintMulti(unprintable) })
}

func TestBuiltinGout(t *testing.T) {
	require.NotPanics(t, func() { MustPrint("foo") })

	require.NotPanics(t, func() { SetWriter(os.Stderr) })
	require.NotPanics(t, func() { SetFormatter(json.Formatter{}) })

	err := Print("foo")
	require.NoError(t, err)

	err = PrintMulti("foo", "bar")
	require.NoError(t, err)

	require.NotPanics(t, func() { MustPrintMulti("foo", "bar") })

	got := Get()
	require.NotNil(t, got)
}

func TestWriterPrinterMulti(t *testing.T) {
	c := New()
	var buf bytes.Buffer
	c.SetWriter(&buf)
	require.NoError(t, c.PrintMulti(
		struct{ Foo string }{Foo: "bar"},
		struct{ Year int }{Year: 1978},
	))
	require.Equal(t, "- foo: bar\n- year: 1978\n", buf.String())
}

func TestWriterAddNewlines(t *testing.T) {
	c := New()
	c.SetFormatter(json.Formatter{})
	var buf bytes.Buffer
	c.SetWriter(&buf)
	c.Print(struct{ Foo string }{Foo: "bar"})
	require.Equal(t, true, strings.HasSuffix(buf.String(), "\n"), "Printer did not have a linebreak suffix")
}
