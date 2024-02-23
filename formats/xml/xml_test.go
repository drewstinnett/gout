package xml

import (
	"encoding/xml"
	"testing"
)

func TestXMLFormatter(t *testing.T) {
	got, err := Formatter{}.Format(struct {
		XMLName xml.Name `xml:"test"`
		Foo     string   `xml:"foo"`
	}{
		Foo: "bar",
	})
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	expect := "<test><foo>bar</foo></test>"
	gotS := string(got)
	if expect != gotS {
		t.Fatalf("expected: %v but got %v", expect, gotS)
	}
}
