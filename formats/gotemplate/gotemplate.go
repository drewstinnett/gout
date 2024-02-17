package gotemplate

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats"
)

type Formatter struct {
	// Template string
	Opts config.FormatterOpts
}

// type TemplateField struct{}
func (w Formatter) Format(v interface{}) ([]byte, error) {
	var tp string
	if t, ok := w.Opts["template"]; !ok {
		tp = `{{ printf "%+v" . }}`
	} else {
		if tp, ok = t.(string); !ok {
			return nil, errors.New("Found a template option, but it's not a string")
		}
	}
	var doc bytes.Buffer
	tmpl, err := template.New("item").Parse(tp)
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
		return &Formatter{}
	})
}
