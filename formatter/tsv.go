package formatter

import (
	"fmt"
	"sort"
)

// TsvFormatter Tab Seperatted Value output.
type tsvFormatter struct{}

// Format How do we actually format YAML?
func (t tsvFormatter) format(data interface{}, config *Config) ([]byte, error) {
	jsonSlice, err := GenericUnmarshal(data)
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
			} else if StringInSlice(k, config.LimitFields) {
				returnString += fmt.Sprint(item[k], "\t")
			}
		}
		returnString += "\n"
	}
	b := []byte(returnString)
	return b, nil
}

// Output Do the output return string here
func (t tsvFormatter) output(data interface{}, config *Config) ([]byte, error) {
	b, nil := t.format(data, config)
	return b, nil
}
