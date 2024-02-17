package gotemplate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGTOFormatterFormat(t *testing.T) {
	f := Formatter{
		Template: `{{ printf "%+v" . }}`,
	}
	got, err := f.Format(struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	})
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "{Title:Ghostbusters Year:1985}", string(got))
}

func TestEmptyTemplate(t *testing.T) {
	f := Formatter{}
	got, err := f.Format(struct {
		Title string
		Year  int
	}{
		Title: "Ghostbusters",
		Year:  1985,
	})
	require.EqualError(t, err, "no Template set for gotemplate")
	require.Nil(t, got)
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
		Template: "{{ .Title }}",
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
		Template: "{{ .NotExistingField }}",
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
		Template: "{{ range . }}{{ .Title }}\n{{ end }}",
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
		Template: "{{ .Title }}",
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
		Template: "{{ .Title }}",
	}
	got, err := f.Format(v)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}
