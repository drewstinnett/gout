package gotemplate

import (
	"bytes"
	"errors"
	"html/template"

	"github.com/drewstinnett/go-output-format/v2/config"
)

type Formatter struct{}

// GoTemplateFormatterOpts Options for using the gotemplate formatter
type FormatterOpts struct {
	Var      interface{}
	Template string
}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return w.FormatWithOpts(v, config.FormatterOpts{
		"template": `{{ . }}`,
	})
}

func (w Formatter) FormatWithOpts(v interface{}, o config.FormatterOpts) ([]byte, error) {
	if _, ok := o["template"]; !ok {
		return nil, errors.New("Must pass 'template' in to options")
	}

	tpl := o["template"].(string)
	var doc bytes.Buffer
	tmpl, err := template.New("item").Parse(tpl)
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&doc, v)
	if err != nil {
		return nil, err
	}
	return doc.Bytes(), nil
}
