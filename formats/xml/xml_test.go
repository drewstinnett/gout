package xml

import (
	"context"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXMLFormatter(t *testing.T) {
	f := Formatter{}
	got, err := f.Format(struct {
		XMLName xml.Name `xml:"test"`
		Foo     string   `xml:"foo"`
	}{
		Foo: "bar",
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("<test><foo>bar</foo></test>"), string(got))
}

func TestXMLFormatterWithContext(t *testing.T) {
	f := Formatter{}
	got, err := f.FormatWithContext(context.Background(), struct {
		XMLName xml.Name `xml:"test"`
		Foo     string   `xml:"foo"`
	}{
		Foo: "bar",
	})
	require.NoError(t, err)
	require.IsType(t, []byte{}, got)
	require.Equal(t, string("<test><foo>bar</foo></test>"), string(got))
}
