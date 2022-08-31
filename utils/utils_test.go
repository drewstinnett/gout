package utils

import (
	"errors"
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
	got, err := GenericUnmarshal(test)
	require.NoError(t, err)
	require.Greater(t, len(got), 0)

	f := fakeValue{
		errors.New("fail"),
	}
	_, err = GenericUnmarshal(f)
	require.Error(t, err)
}

func TestGenericUnmarshalSlice(t *testing.T) {
	test := []map[string]string{
		{"foo": "bar"},
		{"baz": "thing"},
	}
	got, err := GenericUnmarshal(test)
	require.NoError(t, err)
	require.Greater(t, len(got), 0)
}

func TestGenericUnmarshalPtr(t *testing.T) {
	test := []map[string]string{
		{"foo": "bar"},
		{"baz": "thing"},
	}
	got, err := GenericUnmarshal(&test)
	require.NoError(t, err)
	require.Greater(t, len(got), 0)

	f := fakeValue{
		errors.New("fail"),
	}
	_, err = GenericUnmarshal([]fakeValue{f})
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
