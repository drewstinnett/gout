package tsv

import (
	"fmt"
	"sort"

	"github.com/drewstinnett/go-output-format/utils"
)

type Formatter struct {
	LimitFields []string
}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	// func (p *plug) Format(data interface{}, config *config.Config) ([]byte, error) {
	jsonSlice, err := utils.GenericUnmarshal(v)
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
			if len(w.LimitFields) == 0 {
				returnString += fmt.Sprint(item[k], "\t")
			} else if utils.StringInSlice(k, w.LimitFields) {
				returnString += fmt.Sprint(item[k], "\t")
			}
		}
		returnString += "\n"
	}
	b := []byte(returnString)
	return b, nil
}
