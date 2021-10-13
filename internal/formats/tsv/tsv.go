package tsv

import (
	"fmt"
	"sort"

	"github.com/drewstinnett/go-output-format/internal/utils"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

// TsvFormatter Tab Seperatted Value output.
type Formatter struct{}

// Format How do we actually format YAML?
func (t *Formatter) Format(data interface{}, config *config.Config) ([]byte, error) {
	jsonSlice, err := utils.GenericUnmarshal(data)
	if err != nil {
		return nil, err
	}
	returnString := ""

	for _, item := range jsonSlice {
		var keys []string
		for k := range item {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			if len(config.LimitFields) == 0 {
				returnString += fmt.Sprint(item[k], "\t")
			} else if utils.StringInSlice(k, config.LimitFields) {
				returnString += fmt.Sprint(item[k], "\t")
			}
		}
		returnString += "\n"
	}
	b := []byte(returnString)
	return b, nil
}

// Output Do the output return string here
func (t *Formatter) Output(data interface{}, config *config.Config) ([]byte, error) {
	b, nil := t.Format(data, config)
	return b, nil
}

func init() {
	formatter.Add("tsv", func() formatter.Formatter {
		return &Formatter{}
	})
}
