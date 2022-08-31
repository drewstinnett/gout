package gotemplate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGTOFormatterBadInterface(t *testing.T) {
	f := GoTemplateFormatter{}
	got, err := f.Format("Just some string, not a GoTemplateFormatterOpts")
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGTOFormatter(t *testing.T) {
	f := GoTemplateFormatter{}
	opts := GoTemplateFormatterOpts{
		Var: struct {
			Title string
			Year  int
		}{
			Title: "Ghostbusters",
			Year:  1985,
		},
		Template: "{{ .Title }}",
	}
	got, err := f.Format(opts)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters", string(got))
}

func TestGTOFormatterMultiVal(t *testing.T) {
	f := GoTemplateFormatter{}
	opts := GoTemplateFormatterOpts{
		Var: []struct {
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
		},
		Template: "{{ range . }}{{ .Title }}\n{{ end }}",
	}
	got, err := f.Format(opts)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "Ghostbusters\nHalloween\n", string(got))
}
