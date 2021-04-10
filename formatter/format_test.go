package formatter_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestGetFormats(t *testing.T) {
	t.Parallel()
	formats := formatter.GetFormats()
	if !stringInSlice("json", formats) {
		t.Fatalf("GetFormats did not return json")
	}
	if !stringInSlice("yaml", formats) {
		t.Fatalf("GetFormats did not return yaml")
	}
	if !stringInSlice("tsv", formats) {
		t.Fatalf("GetFormats did not return tsv")
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
