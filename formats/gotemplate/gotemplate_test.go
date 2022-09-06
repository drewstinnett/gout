package gotemplate

import (
	"context"
	"testing"

	"github.com/drewstinnett/go-output-format/v2/config"
	"github.com/stretchr/testify/require"
)

func TestGTOFormatterFormat(t *testing.T) {
	f := Formatter{}
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	got, err := f.Format(v)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "{Title:Ghostbusters Year:1985}", string(got))
}

func TestGTOFormatter(t *testing.T) {
	f := Formatter{}
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	opts := config.FormatterOpts{
		"template": "{{ .Title }}",
	}
	got, err := f.formatWithOpts(v, opts)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}

func TestGTOFormatterTemplateError(t *testing.T) {
	f := Formatter{}
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	opts := config.FormatterOpts{
		"template": "{{ .NotExistingField }}",
	}
	got, err := f.formatWithOpts(v, opts)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGTOFormatterMultiVal(t *testing.T) {
	f := Formatter{}
	v := []struct {
		Title string
		Year  int
	}{
		{
			Title: "Ghostbusters",
			Year:  1985,
		},
		{
			Title: "Halloween",
			Year:  1978,
		},
	}
	opts := config.FormatterOpts{
		"template": "{{ range . }}{{ .Title }}\n{{ end }}",
	}
	got, err := f.formatWithOpts(v, opts)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters\nHalloween\n", string(got))
}

func TestGTOWithOptsFormatter(t *testing.T) {
	f := Formatter{}
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	opts := config.FormatterOpts{
		"template": "{{ .Title }}",
	}
	got, err := f.formatWithOpts(v, opts)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}

func TestGTOWithOptsFormatterMissingTemplate(t *testing.T) {
	f := Formatter{}
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	// No 'template' option
	opts := config.FormatterOpts{}
	got, err := f.formatWithOpts(v, opts)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestFormatWithContext(t *testing.T) {
	f := Formatter{}
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	ctx := context.WithValue(context.Background(), TemplateField{}, "{{ .Title }}")
	got, err := f.FormatWithContext(ctx, v)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}
