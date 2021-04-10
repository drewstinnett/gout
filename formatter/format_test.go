package formatter

import (
	"testing"
)

func TestInvalidOutputFormat(t *testing.T) {
	badFormat := "ThisWillNeverBeAValidMarkdown"
	_, got := Formatters[badFormat]
	want := false
	if want != got {
		t.Fatalf(`values not equal ("%t" != "%t")`,
			got,
			want,
		)
	}

	c := &Config{
		Format: badFormat,
	}

	_, err := OutputData("foo", c)
	if err == nil {
		t.Fatalf("OutputData did not err on a bad format")
	}
}

func TestGetFormats(t *testing.T) {
	formats := GetFormats()
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
