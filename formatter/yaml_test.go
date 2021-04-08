package formatter

import "testing"

func TestYamlFormatStruct(t *testing.T) {
	//movie := &Movie{Title: "Halloween", Year: 1978}
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &FormatterConfig{
		Format: "yaml",
	}
	out, _ := OutputData(movie, c)
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
	c := &FormatterConfig{
		Format: "yaml",
	}
	out, _ := OutputData(movies, c)
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
