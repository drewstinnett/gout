package gotemplate

import (
	"testing"

	"github.com/drewstinnett/gout/v2/config"
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
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	f := Formatter{
		Opts: map[string]interface{}{
			"template": "{{ .Title }}",
		},
	}
	got, err := f.Format(v)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}

func TestGTOFormatterTemplateError(t *testing.T) {
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	f := Formatter{
		Opts: map[string]interface{}{
			"template": "{{ .NotExistingField }}",
		},
	}
	got, err := f.Format(v)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGTOFormatterMultiVal(t *testing.T) {
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
	f := Formatter{
		Opts: map[string]interface{}{
			"template": "{{ range . }}{{ .Title }}\n{{ end }}",
		},
	}
	got, err := f.Format(v)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters\nHalloween\n", string(got))
}

func TestGTOWithOptsFormatter(t *testing.T) {
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	f := Formatter{
		Opts: config.FormatterOpts{
			"template": "{{ .Title }}",
		},
	}
	got, err := f.Format(v)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}

/*
func TestGTOWithOptsFormatterMissingTemplate(t *testing.T) {
	f := Formatter{
		Template: "",
	}
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	// No 'template' option
	got, err := f.Format(v)
	require.Error(t, err)
	require.Nil(t, got)
}
*/

func TestFormatWithContext(t *testing.T) {
	v := struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	}
	f := Formatter{
		Opts: config.FormatterOpts{
			"template": "{{ .Title }}",
		},
	}
	got, err := f.Format(v)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}
