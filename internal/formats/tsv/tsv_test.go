package tsv_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"github.com/stretchr/testify/require"
)

func TestTSVInvalidDataType(t *testing.T) {
	c := &config.Config{
		Format: "tsv",
	}
	_, err := formatter.OutputData(func() {}, c)
	require.Error(t, err)
}

func TestTSVInvalidDataStruct(t *testing.T) {
	c := &config.Config{
		Format: "tsv",
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

func TestTSVInvalidDataSlice(t *testing.T) {
	c := &config.Config{
		Format: "tsv",
	}
	movies := []struct {
		Title fakeValue
		Year  int
	}{
		{
			fakeValue{errors.New("fail_the_movie")},
			1984,
		},
		{
			fakeValue{errors.New("fail_the_movie_again")},
			1985,
		},
	}
	_, err := formatter.OutputData(movies, c)
	require.Error(t, err)
}

func TestTSVField(t *testing.T) {
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &config.Config{
		Format:      "tsv",
		LimitFields: []string{"Title"},
	}
	out, _ := formatter.OutputData(&movie, c)
	got := strings.TrimSpace(string(out))

	want := "Halloween"
	require.Equal(t, want, got)
}

func TestTSVFormatStructPtr(t *testing.T) {
	t.Parallel()
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &config.Config{
		Format: "tsv",
	}
	out, _ := formatter.OutputData(&movie, c)
	got := strings.TrimSpace(string(out))

	want := "Halloween\t1978"
	require.Equal(t, want, got)
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
	c := &config.Config{
		Format: "tsv",
	}
	out, _ := formatter.OutputData(movie, c)
	got := strings.TrimSpace(string(out))

	want := "Halloween\t1978"
	require.Equal(t, want, got)
}

func TestTSVFormatStructListPtr(t *testing.T) {
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
		Format: "tsv",
	}
	out, _ := formatter.OutputData(&movies, c)
	got := strings.Replace(strings.TrimSpace(string(out)), "\t", " ", -1)

	require.Contains(t, got, "Halloween 1978")
	require.Contains(t, got, "Phantasm 1979")
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
	c := &config.Config{
		Format: "tsv",
	}
	out, _ := formatter.OutputData(movies, c)
	got := strings.Replace(strings.TrimSpace(string(out)), "\t", " ", -1)

	require.Contains(t, got, "Halloween 1978")
	require.Contains(t, got, "Phantasm 1979")
}

type fakeValue struct {
	err error
}
