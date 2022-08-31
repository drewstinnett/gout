package gotemplate

import (
	"bytes"
	"errors"
	"html/template"
)

type GoTemplateFormatter struct{}

// GoTemplateFormatterOpts Options for using the gotemplate formatter
type GoTemplateFormatterOpts struct {
	Var      interface{}
	Template string
}

func (w GoTemplateFormatter) Format(v interface{}) ([]byte, error) {
	opts, ok := v.(GoTemplateFormatterOpts)
	if !ok {
		return nil, errors.New("Must pass in a GoTemplateFormatterOpts")
	}

	var doc bytes.Buffer
	tmpl, err := template.New("item").Parse(opts.Template)
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&doc, opts.Var)
	if err != nil {
		return nil, err
	}
	return doc.Bytes(), nil
}
