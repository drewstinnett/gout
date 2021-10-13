package yaml_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"gotest.tools/assert"
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
	assert.Equal(t, want, got)
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
	assert.Equal(t, want, got)
}
