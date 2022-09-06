package gotemplate

import (
	"bytes"
	"context"
	"errors"
	"html/template"

	"github.com/drewstinnett/go-output-format/v2/config"
)

type Formatter struct {
	ctx context.Context
}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	var t any
	if w.ctx != nil {
		t = w.ctx.Value("template")
	}
	if t == nil {
		t = `{{ printf "%+v" . }}`
	}
	return w.FormatWithOpts(v, config.FormatterOpts{
		"template": t,
	})
}

func (w Formatter) FormatWithContext(ctx context.Context, v interface{}) ([]byte, error) {
	return w.withContext(ctx).Format(v)
}

func (w *Formatter) withContext(ctx context.Context) *Formatter {
	w.ctx = ctx
	return w
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
