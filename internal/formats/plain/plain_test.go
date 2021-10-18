package plain_test

import (
	"testing"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"github.com/stretchr/testify/require"
)

func TestPlainFormatStruct(t *testing.T) {
	t.Parallel()
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &config.Config{
		Format: "plain",
	}
	out, _ := formatter.OutputData(movie, c)
	got := string(out)

	want := "{Title:Halloween Year:1978}"
	require.Equal(t, want, got)
}

func TestPlainFormatStructList(t *testing.T) {
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
		Format: "plain",
	}
	out, _ := formatter.OutputData(movies, c)
	got := string(out)

	want := "[{Title:Halloween Year:1978} {Title:Phantasm Year:1979}]"
	require.Equal(t, want, got)
}
