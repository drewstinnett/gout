/*
Package gotemplate provides a way to template output as a plugin to Gout
*/
package gotemplate

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/drewstinnett/gout/v2/formats"
)

// Formatter is the base object for the gotemplate object
type Formatter struct {
	Template string
	// Opts config.FormatterOpts
}

// Format satisfies the formats.Format interface
func (w Formatter) Format(v any) ([]byte, error) {
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
