package formatter_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestSliceContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := formatter.StringInSlice("bar", s)
	if got != true {
		t.Fatalf("Could not find 'bar' in slice")
	}
}

func TestSliceNotContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := formatter.StringInSlice("NeverExists", s)
	if got != false {
		t.Fatalf("Found something nit shouldn't have")
	}
}
