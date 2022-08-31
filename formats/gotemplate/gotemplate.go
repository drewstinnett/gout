package gotemplate

import (
	"bytes"
	"errors"
	"html/template"
)

type Formatter struct{}

// GoTemplateFormatterOpts Options for using the gotemplate formatter
type FormatterOpts struct {
	Var      interface{}
	Template string
}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	opts, ok := v.(FormatterOpts)
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
