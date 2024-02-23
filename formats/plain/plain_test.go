package plain

import (
	"testing"
)

func TestPlainFormatter(t *testing.T) {
	got, err := Formatter{}.Format(struct {
		Foo string
	}{
		Foo: "bar",
	})
	if err != nil {
		t.Fatalf("got an unexpected error: %v", err)
	}
	expect := "{Foo:bar}\n"
	gotS := string(got)
	if gotS != expect {
		t.Fatalf("expected: %v but got %v", expect, got)
	}
}
