package formatter_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/drewstinnett/go-output-format/formatter"
)

func TestGoTemplateInvalidDataType(t *testing.T) {
	c := &formatter.Config{
		Format:   "gotemplate",
		Template: "{{ .Title }}",
	}
	_, err := formatter.OutputData(func() {}, c)
	if err == nil {
		t.Fatalf(`Did not return an error on bad data input`)
	}
}

func TestGoTemplateMissingTemplate(t *testing.T) {
	c := &formatter.Config{
		Format: "gotemplate",
	}
	_, err := formatter.OutputData("foo", c)
	if err == nil {
		t.Fatal("Did not pass a Template parameter to gotemplate, but didn't return an error")
	}
}

func TestGoTemplateInvalidTemplate(t *testing.T) {
	c := &formatter.Config{
		Format:   "gotemplate",
		Template: "{{ .Name ",
	}
	_, err := formatter.OutputData("foo", c)
	if err == nil {
		t.Fatal("Passed a bad template, but didn't error")
	}
}

func TestGoTemplateInvalidDataStruct(t *testing.T) {
	c := &formatter.Config{
		Format:   "gotemplate",
		Template: "{{ .Title }}",
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

func TestGoTemplateFormatStruct(t *testing.T) {
	t.Parallel()
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &formatter.Config{
		Format:   "gotemplate",
		Template: "{{ .Title }}",
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

func TestGoTemplateFormatStructList(t *testing.T) {
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
		Format:   "gotemplate",
		Template: "{{ .Title }}\n",
	}
	out, _ := formatter.OutputData(&movies, c)
	got := strings.TrimSpace(string(out))
	want := "Halloween\nPhantasm"

	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}
