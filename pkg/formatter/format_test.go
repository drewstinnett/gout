package formatter_test

import (
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
