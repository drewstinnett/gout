package formatter_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestYamlFormatStruct(t *testing.T) {
	//movie := &Movie{Title: "Halloween", Year: 1978}
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &formatter.Config{
		Format: "yaml",
	}
	out, _ := formatter.OutputData(movie, c)
	got := string(out)

	want := `title: Halloween
year: 1978
`
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}

func TestYamlFormatStructList(t *testing.T) {
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
		Format: "yaml",
	}
	out, _ := formatter.OutputData(movies, c)
	got := string(out)

	want := `- title: Halloween
  year: 1978
- title: Phantasm
  year: 1979
`
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}
