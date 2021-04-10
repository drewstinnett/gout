package formatter_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestJSONFormatStruct(t *testing.T) {
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &formatter.Config{
		Format: "json",
	}
	out, _ := formatter.OutputData(movie, c)
	got := string(out)

	want := `{
  "Title": "Halloween",
  "Year": 1978
}`
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}

func TestJsonFormatStructList(t *testing.T) {
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
		Format: "json",
	}
	out, _ := formatter.OutputData(movies, c)
	got := string(out)

	want := `[
  {
    "Title": "Halloween",
    "Year": 1978
  },
  {
    "Title": "Phantasm",
    "Year": 1979
  }
]`
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}
