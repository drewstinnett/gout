package formats_test

import (
	"slices"
	"testing"

	"github.com/drewstinnett/gout/v2/formats"
	_ "github.com/drewstinnett/gout/v2/formats/builtin"
)

func TestBuiltinRegistry(t *testing.T) {
	got := formats.Names()
	formats := []string{"gotemplate", "json", "plain", "xml", "yaml"}
	for _, expect := range formats {
		if len(got) != len(formats) {
			t.Fatalf("expected len to be %v but got %v", len(formats), len(got))
		}
		if !slices.Contains(got, expect) {
			t.Fatalf("expected %v to be in %v", expect, got)
		}
	}
}
