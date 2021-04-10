package formatter_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestInvalidOutputFormat(t *testing.T) {
	t.Parallel()
	badFormat := "ThisWillNeverBeAValidMarkdown"
	_, got := formatter.Formatters[badFormat]
	want := false
	if want != got {
		t.Fatalf(`values not equal ("%t" != "%t")`,
			got,
			want,
		)
	}

	c := &formatter.Config{
		Format: badFormat,
	}

	_, err := formatter.OutputData("foo", c)
	if err == nil {
		t.Fatalf("OutputData did not err on a bad format")
	}
}

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
