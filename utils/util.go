package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// GenericUnmarshal Given an arbitrary piece of data, return a slice of json data
func GenericUnmarshal(data interface{}) ([]map[string]interface{}, error) {
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
		return nil, fmt.Errorf("Unknown type of data for gotemplate: %s", objType)
	}
	return jsonSlice, nil
}

// StringInSlice Check if a slice for a string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
