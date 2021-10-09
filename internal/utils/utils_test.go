package utils_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/internal/utils"
)

func TestSliceContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := utils.StringInSlice("bar", s)
	if got != true {
		t.Fatalf("Could not find 'bar' in slice")
	}
}

func TestSliceNotContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := utils.StringInSlice("NeverExists", s)
	if got != false {
		t.Fatalf("Found something nit shouldn't have")
	}
}
