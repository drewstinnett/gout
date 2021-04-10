package formatter

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
)

// TsvFormatter Tab Seperatted Value output
type TsvFormatter struct{}

// Format How do we actually format YAML?
func (y TsvFormatter) Format(data interface{}) ([]byte, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var jsonSlice []map[string]interface{}
	switch objType := reflect.ValueOf(data).Elem().Kind(); objType {
	case reflect.Struct:
		jsonMap := make(map[string]interface{})
		err := json.Unmarshal(j, &jsonMap)
		if err != nil {
			return nil, err
		}
		jsonSlice = append(jsonSlice, jsonMap)
	case reflect.Slice:
		err := json.Unmarshal(j, &jsonSlice)
		if err != nil {
			return nil, err
		}
	default:
		log.Println("ERR: ", objType)
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
			returnString += fmt.Sprint(item[k], "\t")
		}
		returnString += fmt.Sprint("\n")
	}
	b := []byte(returnString)
	return b, nil
}

// Output Do the output return string here
func (y TsvFormatter) Output(data interface{}) ([]byte, error) {
	b, err := y.Format(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
