package formats_test

import (
	"errors"
	"testing"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"github.com/stretchr/testify/require"
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
	c := &config.Config{
		Format: "json",
	}
	out, _ := formatter.OutputData(movie, c)
	got := string(out)

	want := `{
  "Title": "Halloween",
  "Year": 1978
}`
	require.Equal(t, want, got)
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
	c := &config.Config{
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
	require.Equal(t, want, got)
}

func TestJSONInvalidDataStruct(t *testing.T) {
	c := &config.Config{
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
	require.Error(t, err)
}
