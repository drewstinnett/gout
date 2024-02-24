package json

import (
	"testing"

	"github.com/drewstinnett/gout/v2/formats"
)

func TestJSONFormatter(t *testing.T) {
	for desc, tt := range map[string]struct {
		given  formats.Formatter
		expect string
	}{
		"default": {
			given:  Formatter{},
			expect: `{"Foo":"bar"}`,
		},
		"indented": {
			given: Formatter{
				indent: true,
			},
			expect: "{\n  \"Foo\": \"bar\"\n}",
		},
	} {
		got, err := tt.given.Format(struct{ Foo string }{Foo: "bar"})
		if err != nil {
			t.Fatalf("%v: got an unexpected error: %v", desc, err)
		}
		gotS := string(got)
		if tt.expect != gotS {
			t.Fatalf("%v: expected: %v, but got: %v", desc, tt.expect, gotS)
		}
	}
}
