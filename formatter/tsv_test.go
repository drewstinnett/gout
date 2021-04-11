package formatter_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestTSVInvalidDataType(t *testing.T) {
	c := &formatter.Config{
		Format: "tsv",
	}
	_, err := formatter.OutputData(func() {}, c)
	if err == nil {
		t.Fatalf(`Did not return an error on bad data input`)
	}
}

func TestTSVInvalidDataStruct(t *testing.T) {
	c := &formatter.Config{
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
	if err == nil {
		t.Fatalf("Did not return on bad data struct")
	}
}

func TestTSVInvalidDataSlice(t *testing.T) {
	c := &formatter.Config{
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
	if err == nil {
		t.Fatalf("Did not return on bad data slice")
	}
}

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

func TestTSVFormatStructPtr(t *testing.T) {
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
	out, _ := formatter.OutputData(movie, c)
	got := strings.TrimSpace(string(out))

	want := "Halloween\t1978"
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
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
	out, _ := formatter.OutputData(movies, c)
	got := strings.Replace(strings.TrimSpace(string(out)), "\t", " ", -1)

	if !strings.Contains(got, "Halloween 1978") {
		t.Fatalf(`%s does not contain "Halloween 1978"`, got)
	}
	if !strings.Contains(got, "Phantasm 1979") {
		t.Fatalf(`%s does not contain "Phantasm 1979"`, got)
	}
}
