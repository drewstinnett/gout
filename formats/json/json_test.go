package json

import (
	"testing"
)

func TestJSONFormatter(t *testing.T) {
	f := Formatter{}
	got, err := f.Format(struct {
		Foo string
	}{
		Foo: "bar",
	})
	if err != nil {
		t.Fatalf("got an unexpected error: %v", err)
	}
	expect := `{"Foo":"bar"}`
	gotS := string(got)
	if expect != gotS {
		t.Fatalf("expected: %v, but got: %v", expect, got)
	}
}
