package formats_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

func TestYamlFormatStruct(t *testing.T) {
	t.Parallel()
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &config.Config{
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
	c := &config.Config{
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
