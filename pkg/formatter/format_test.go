package formatter_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/internal/utils"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

func TestBadFormat(t *testing.T) {
	c := &config.Config{
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
	if !utils.StringInSlice("json", formats) {
		t.Fatalf("GetFormats did not return json")
	}
	if !utils.StringInSlice("yaml", formats) {
		t.Fatalf("GetFormats did not return yaml")
	}
	if !utils.StringInSlice("tsv", formats) {
		t.Fatalf("GetFormats did not return tsv")
	}
}
