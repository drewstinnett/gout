package gotemplate

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/drewstinnett/go-output-format/internal/utils"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

// GotemplateFormatter Tab Seperatted Value output.
type GoTemplateFormatter struct{}

// Format How do we actually format the data back?
func (g *GoTemplateFormatter) Format(data interface{}, config *config.Config) ([]byte, error) {
	if config.Template == "" {
		return nil, fmt.Errorf("Missing required config value of 'Template' for gotemplate")
	}
	jsonSlice, err := utils.GenericUnmarshal(data)
	if err != nil {
		return nil, err
	}
	returnString := ""

	for _, item := range jsonSlice {
		var doc bytes.Buffer
		tmpl, err := template.New("item").Parse(config.Template)
		if err != nil {
			return nil, err
		}
		err = tmpl.Execute(&doc, item)
		if err != nil {
			return nil, err
		}
		returnString += doc.String()
	}
	b := []byte(returnString)
	return b, nil
}

// Output Do the output return string here
func (g *GoTemplateFormatter) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, nil := g.Format(data, config)
	return b, nil
}

func init() {
	formatter.Add("gotemplate", func() formatter.Formatter {
		return &GoTemplateFormatter{}
	})
}
