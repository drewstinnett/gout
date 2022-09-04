package main

import (
	"fmt"

	"github.com/drewstinnett/go-output-format/v2/formats/gotemplate"
	"github.com/drewstinnett/go-output-format/v2/gout"
)

type sample struct {
	FirstName string
	LastName  string
	Age       int
}

type sampleList []sample

func main() {
	people := &sampleList{
		sample{FirstName: "Jason", LastName: "Vorhees", Age: 11},
		sample{FirstName: "Freddy", LastName: "Krueger", Age: 35},
		sample{FirstName: "Michael", LastName: "Myers", Age: 13},
	}
	c, _ := gout.New()
	c.SetFormatter(gotemplate.Formatter{})
	fmt.Printf("# Format: gotemplate\n")
	fmt.Println("## People")
	c.MustPrint(gotemplate.FormatterOpts{
		Var:      people,
		Template: "{{ range . }}{{ .FirstName }} {{ .LastName }} is {{ .Age }} years old\n{{ end }}",
	})
	fmt.Println()

	// fmt.Println(string(b))
}
