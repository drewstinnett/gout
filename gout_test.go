package gout

import (
	"bytes"
	"strings"
	"testing"

	"github.com/drewstinnett/gout/v2/formats/json"
	"github.com/drewstinnett/gout/v2/formats/plain"
)

func TestNewWriter(t *testing.T) {
	if got := New(); got == nil {
		t.Fatal("got a nil client with New()")
	}
}

func TestSetGoutWriter(t *testing.T) {
	c := New()
	var buf bytes.Buffer
	c.SetWriter(&buf)
	c.Print(struct{ Foo string }{Foo: "bar"})
	expect := "foo: bar\n"
	got := buf.String()
	if expect != got {
		t.Fatalf("expected: %v but got %v", expect, got)
	}

	// Make sure we can change the formatter
	c.SetFormatter(plain.Formatter{})
	buf = bytes.Buffer{}
	c.SetFormatterString("json")
	c.Print(struct{ Foo string }{Foo: "bar"})
	expect = "{\"Foo\":\"bar\"}\n"
	got = buf.String()
	if expect != got {
		t.Fatalf("expected %v but got %v", expect, got)
	}
}

func TestNewGoutWithWriter(t *testing.T) {
	var b bytes.Buffer
	New(WithWriter(&b)).Print(struct{ Foo string }{Foo: "bar"})
	expect := "foo: bar\n"
	got := b.String()
	if expect != got {
		t.Fatalf("expected %v but got %v", expect, got)
	}
}

func TestNewGoutWithFormatter(t *testing.T) {
	var b bytes.Buffer
	New(WithWriter(&b), WithFormatter(plain.Formatter{})).Print(struct{ Foo string }{Foo: "bar"})
	expect := "{Foo:bar}\n"
	got := b.String()
	if expect != got {
		t.Fatalf("expected %v but got %v", expect, got)
	}
}

func TestPrintError(t *testing.T) {
	c := New()
	expectErr := "panic while attempting to format: cannot marshal type: chan int"
	if err := c.Print(make(chan int)); err.Error() != expectErr {
		t.Fatalf("expected error to be:\n%v\n\nbut got:\n%v\n\n", expectErr, err.Error())
	}
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestPrintPanic(t *testing.T) {
	c := New()
	unprintable := make(chan int)
	assertPanic(t, func() { c.MustPrint(unprintable) })
	assertPanic(t, func() { c.MustPrintMulti(unprintable) })
}

func TestBuiltinGout(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	MustPrint("foo")
	SetFormatter(json.Formatter{})

	expect := "foo\n"
	got := b.String()
	if expect != got {
		t.Fatalf("expected:\n%v\n\nbut got:\n%v\n\n", expect, got)
	}
	if err := Print("foo"); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if err := PrintMulti("foo", "bar"); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	MustPrintMulti("foo", "bar")

	if Get() == nil {
		t.Fatalf("got nil from Get()")
	}

	if err := SetFormatterString("plain"); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}

	expectErr := "unknown formatter name: never-exists"
	if err := SetFormatterString("never-exists"); err.Error() != expectErr {
		t.Fatalf("expected error '%v', but got '%v'", expectErr, err)
	}
}

func TestWriterPrinterMulti(t *testing.T) {
	c := New()
	var buf bytes.Buffer
	c.SetWriter(&buf)
	if err := c.PrintMulti(
		struct{ Foo string }{Foo: "bar"},
		struct{ Year int }{Year: 1978},
	); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	expect := "- foo: bar\n- year: 1978\n"
	got := buf.String()
	if expect != got {
		t.Fatalf("expected %v but got %v", expect, got)
	}
}

func TestWriterAddNewlines(t *testing.T) {
	c := New()
	c.SetFormatter(json.Formatter{})
	var buf bytes.Buffer
	c.SetWriter(&buf)
	c.Print(struct{ Foo string }{Foo: "bar"})
	if !strings.HasSuffix(buf.String(), "\n") {
		t.Fatal("Print did not have a linebreak suffix")
	}
}
