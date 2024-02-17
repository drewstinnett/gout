package main

import (
	"fmt"

	gout "github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats/gotemplate"
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
	g := gout.New()
	g.SetFormatter(gotemplate.Formatter{
		Opts: config.FormatterOpts{
			"template": "{{ range . }}{{ .FirstName }} {{ .LastName }} is {{ .Age }} years old\n{{ end }}",
		},
	})
	//"template": "{{ range . }}{{ .FirstName }} {{ .LastName }} is {{ .Age }} years old\n{{ end }}",
	fmt.Printf("# Format: gotemplate\n## People\n")
	g.MustPrint(people)
	fmt.Println()

	// fmt.Println(string(b))
}
