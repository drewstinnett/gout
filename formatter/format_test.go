package formatter

import (
	"testing"
)

func TestInvalidOutputFormat(t *testing.T) {
	badFormat := "ThisWillNeverBeAValidMarkdown"
	_, got := Formatters[badFormat]
	want := false
	if want != got {
		t.Fatalf(`values not equal ("%t" != "%t")`,
			got,
			want,
		)
	}
}
