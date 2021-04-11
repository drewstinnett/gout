package formatter

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"

	"github.com/drewstinnett/go-output-format/helpers"
)

// TsvFormatter Tab Seperatted Value output.
type tsvFormatter struct{}

// Format How do we actually format YAML?
func (t tsvFormatter) format(data interface{}, config *Config) ([]byte, error) {
	j, _ := json.Marshal(data)
	var jsonSlice []map[string]interface{}
	baseObjType := reflect.ValueOf(data).Kind()
	var objType string
	if baseObjType == reflect.Struct {
		objType = "struct"
	} else if baseObjType == reflect.Slice {
		objType = "slice"
	} else if baseObjType == reflect.Ptr {
		objType = reflect.ValueOf(data).Elem().Kind().String()
	}
	log.Println("OBJ Type is: ", objType)
	switch objType {
	case "struct":
		jsonMap := make(map[string]interface{})
		err := json.Unmarshal(j, &jsonMap)
		if err != nil {
			return nil, err
		}
		jsonSlice = append(jsonSlice, jsonMap)
	case "slice":
		err := json.Unmarshal(j, &jsonSlice)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unknown type of data for tsv: %s", objType)
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
			} else if helpers.StringInSlice(k, config.LimitFields) {
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
