package formatter_test

import (
	"errors"
	"testing"

	_ "github.com/drewstinnett/go-output-format/internal/formats/all"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
	"github.com/stretchr/testify/require"
)

func TestBadFormat(t *testing.T) {
	c := &config.Config{
		Format: "NeverExist",
	}
	_, err := formatter.OutputData(nil, c)
	require.Error(t, err)
}

func TestGetFormats(t *testing.T) {
	t.Parallel()
	formats := formatter.GetFormats()
	require.Subset(t, formats, []string{"gotemplate", "yaml"})
}

func TestOutputFormats(t *testing.T) {
	c := &config.Config{
		Format: "json",
	}
	out, err := formatter.OutputData("hi", c)
	require.NoError(t, err)
	require.Equal(t, "\"hi\"", string(out))
}

func TestOutputError(t *testing.T) {
	c := &config.Config{
		Format: "json",
	}
	_, err := formatter.OutputData(fakeValue{
		errors.New("err"),
	}, c)
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
