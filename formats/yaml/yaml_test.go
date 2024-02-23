package yaml

import (
	"testing"
)

func TestYAMLFormatter(t *testing.T) {
	f := Formatter{}
	got, err := f.Format(struct {
		Foo string
	}{
		Foo: "bar",
	})
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	expect := "foo: bar\n"
	gotS := string(got)
	if expect != gotS {
		t.Fatalf("expected %v but got %v", expect, got)
	}
}
