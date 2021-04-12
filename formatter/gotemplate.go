package formatter

import (
	"bytes"
	"fmt"
	"text/template"
)

// GotemplateFormatter Tab Seperatted Value output.
type gotemplateFormatter struct{}

// Format How do we actually format the data back?
func (g gotemplateFormatter) format(data interface{}, config *Config) ([]byte, error) {
	if config.Template == "" {
		return nil, fmt.Errorf("Missing required config value of 'Template' for gotemplate")
	}
	jsonSlice, err := GenericUnmarshal(data)
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
func (g gotemplateFormatter) output(data interface{}, config *Config) ([]byte, error) {
	b, nil := g.format(data, config)
	return b, nil
}
