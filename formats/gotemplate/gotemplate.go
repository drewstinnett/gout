package gotemplate

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/drewstinnett/gout/v2/formats"
)

type Formatter struct {
	Template string
	// Opts config.FormatterOpts
}

// type TemplateField struct{}
func (w Formatter) Format(v interface{}) ([]byte, error) {
	if w.Template == "" {
		return nil, errors.New("no Template set for gotemplate")
	}
	var doc bytes.Buffer
	tmpl, err := template.New("item").Parse(w.Template)
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&doc, v)
	if err != nil {
		return nil, err
	}
	return doc.Bytes(), nil
}

func init() {
	formats.Add("gotemplate", func() formats.Formatter {
		return &Formatter{
			Template: `{{ printf "%+v" . }}`,
		}
	})
}
