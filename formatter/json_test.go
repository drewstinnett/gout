package formatter_test

import (
	"errors"
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestJSONFormatStruct(t *testing.T) {
	t.Parallel()
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

func TestJSONInvalidDataStruct(t *testing.T) {
	c := &formatter.Config{
		Format: "json",
	}
	movie := struct {
		Title fakeValue
		Year  int
	}{
		fakeValue{errors.New("fail_the_movie")},
		1984,
	}
	_, err := formatter.OutputData(movie, c)
	if err == nil {
		t.Fatalf("Did not return on bad JSON data struct")
	}
}
