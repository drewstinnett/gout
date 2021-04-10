package formatter_test

import (
	"strings"
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestTSVField(t *testing.T) {
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &formatter.Config{
		Format:      "tsv",
		LimitFields: []string{"Title"},
	}
	out, _ := formatter.OutputData(&movie, c)
	got := strings.TrimSpace(string(out))

	want := "Halloween"
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}
func TestTSVFormatStruct(t *testing.T) {
	t.Parallel()
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &formatter.Config{
		Format: "tsv",
	}
	out, _ := formatter.OutputData(&movie, c)
	got := strings.TrimSpace(string(out))

	want := "Halloween\t1978"
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}

func TestTSVFormatStructList(t *testing.T) {
	t.Parallel()
	movies := []struct {
		Title string
		Year  int
	}{
		{
			"Halloween",
			1978,
		},
		{
			"Phantasm",
			1979,
		},
	}
	c := &formatter.Config{
		Format: "tsv",
	}
	out, _ := formatter.OutputData(&movies, c)
	got := strings.Replace(strings.TrimSpace(string(out)), "\t", " ", -1)

	if !strings.Contains(got, "Halloween 1978") {
		t.Fatalf(`%s does not contain "Halloween 1978"`, got)
	}
	if !strings.Contains(got, "Phantasm 1979") {
		t.Fatalf(`%s does not contain "Phantasm 1979"`, got)
	}
}
