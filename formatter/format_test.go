package formatter_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
	"github.com/drewstinnett/go-output-format/helpers"
)

func TestBadFormat(t *testing.T) {
	c := &formatter.Config{
		Format: "NeverExist",
	}
	_, err := formatter.OutputData(nil, c)
	if err == nil {
		t.Fatalf("Using a bad Format did not cause an error")
	}

}

func TestGetFormats(t *testing.T) {
	t.Parallel()
	formats := formatter.GetFormats()
	if !helpers.StringInSlice("json", formats) {
		t.Fatalf("GetFormats did not return json")
	}
	if !helpers.StringInSlice("yaml", formats) {
		t.Fatalf("GetFormats did not return yaml")
	}
	if !helpers.StringInSlice("tsv", formats) {
		t.Fatalf("GetFormats did not return tsv")
	}
}
