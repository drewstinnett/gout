package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := StringInSlice("bar", s)
	require.True(t, got)
}

func TestSliceNotContains(t *testing.T) {
	t.Parallel()
	s := []string{"foo", "bar", "baz"}

	got := StringInSlice("NeverExists", s)
	require.False(t, got)
}

func TestGenericUnmarshalStruct(t *testing.T) {
	type foo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	test := foo{
		Name: "Drew",
		Age:  25,
	}
	got, err := genericUnmarshal(test)
	require.NoError(t, err)
	require.Greater(t, len(got), 0)

	f := fakeValue{
		errors.New("fail"),
	}
	_, err = genericUnmarshal(f)
	require.Error(t, err)
}

func TestGenericUnmarshalSlice(t *testing.T) {
	test := []map[string]string{
		{"foo": "bar"},
		{"baz": "thing"},
	}
	got, err := genericUnmarshal(test)
	require.NoError(t, err)
	require.Greater(t, len(got), 0)
}

func TestGenericUnmarshalPtr(t *testing.T) {
	test := []map[string]string{
		{"foo": "bar"},
		{"baz": "thing"},
	}
	got, err := genericUnmarshal(&test)
	require.NoError(t, err)
	require.Greater(t, len(got), 0)

	f := fakeValue{
		errors.New("fail"),
	}
	_, err = genericUnmarshal([]fakeValue{f})
	require.Error(t, err)
}

type fakeValue struct {
	err error
}

func (v fakeValue) MarshalJSON() ([]byte, error) {
	if v.err != nil {
		return nil, v.err
	}

	return []byte(`null`), v.err
}

// GenericUnmarshal Given an arbitrary piece of data, return a slice of json data
func genericUnmarshal(data interface{}) ([]map[string]interface{}, error) {
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
