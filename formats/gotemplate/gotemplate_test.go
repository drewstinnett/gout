package gotemplate

import (
	"testing"
)

func TestGTOFormatterFormat(t *testing.T) {
	type movie struct {
		Title string
		Year  int
	}

	tests := map[string]struct {
		given     []movie
		givenT    string
		expect    string
		expectErr string
	}{
		"custom-template": {
			given: []movie{
				{Title: "Ghostbusters", Year: 1985},
			},
			givenT: `{{ printf "%+v" . }}`,
			expect: "{Title:Ghostbusters Year:1985}",
		},
		"missing-template": {
			given: []movie{
				{Title: "Ghostbusters"},
			},
			expectErr: "no Template set for gotemplate",
		},
		"bad-template": {
			given: []movie{
				{Title: "Ghostbusters"},
			},
			givenT:    "{{ .NotExistingField }}",
			expectErr: `template: item:1:3: executing "item" at <.NotExistingField>: can't evaluate field NotExistingField in type gotemplate.movie`,
		},
		"multiple-items": {
			given: []movie{
				{Title: "Ghostbusters"},
				{Title: "Halloween"},
			},
			givenT: "{{ range . }}{{ .Title }}\n{{ end }}",
			expect: "Ghostbusters\nHalloween\n",
		},
	}
	for desc, tt := range tests {
		// If not multiple 'givens', only pass in a single non-iteralbe interface
		var realGiven any
		if len(tt.given) == 1 {
			realGiven = tt.given[0]
		} else {
			realGiven = tt.given
		}
		got, err := Formatter{
			Template: tt.givenT,
		}.Format(realGiven)
		if tt.expectErr != "" {
			if err.Error() != tt.expectErr {
				t.Fatalf("%v: expected an error of:\n'%v'\nbut got\n'%v", desc, tt.expectErr, err.Error())
			}
		} else {
			if err != nil {
				t.Fatalf("%v: got an error when none was expected: %v", desc, err)
			}
			if tt.expect != string(got) {
				t.Fatalf("%v: expected %v but got %v", desc, tt.expect, string(got))
			}
		}
	}
}
