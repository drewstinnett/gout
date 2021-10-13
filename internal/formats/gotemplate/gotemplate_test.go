package gotemplate_test

import (
	"errors"
	"strings"
	"testing"

	_ "github.com/drewstinnett/go-output-format/internal/formats/gotemplate"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"github.com/stretchr/testify/require"
)

func TestGoTemplateInvalidDataType(t *testing.T) {
	c := &config.Config{
		Format:   "gotemplate",
		Template: "{{ .Title }}",
	}
	_, err := formatter.OutputData(func() {}, c)
	require.Error(t, err)
}

func TestGoTemplateMissingTemplate(t *testing.T) {
	c := &config.Config{
		Format: "gotemplate",
	}
	_, err := formatter.OutputData("foo", c)
	require.Error(t, err)
}

func TestGoTemplateInvalidTemplate(t *testing.T) {
	c := &config.Config{
		Format:   "gotemplate",
		Template: "{{ .Name ",
	}
	_, err := formatter.OutputData("foo", c)
	require.Error(t, err)
}

func TestGoTemplateInvalidDataStruct(t *testing.T) {
	c := &config.Config{
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
	require.Error(t, err)
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
	c := &config.Config{
		Format:   "gotemplate",
		Template: "{{ .Title }}",
	}
	out, _ := formatter.OutputData(&movie, c)
	got := strings.TrimSpace(string(out))

	want := "Halloween"
	require.Equal(t, want, got)
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
	c := &config.Config{
		Format:   "gotemplate",
		Template: "{{ .Title }}\n",
	}
	out, _ := formatter.OutputData(&movies, c)
	got := strings.TrimSpace(string(out))
	want := "Halloween\nPhantasm"

	require.Equal(t, want, got)
}

type fakeValue struct {
	err error
}

func (v fakeValue) MarshalJSON() ([]byte, error) {
	if v.err != nil {
		return nil, v.err
	}

	return []byte(`null`), v.err
}
