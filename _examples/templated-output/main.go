package main

import (
	"fmt"
	"log"

	_ "github.com/drewstinnett/go-output-format/internal/formats/gotemplate"
	"github.com/drewstinnett/go-output-format/pkg/config"
	"github.com/drewstinnett/go-output-format/pkg/formatter"
)

type sample struct {
	FirstName string
	LastName  string
	Age       int
}

type sampleList []sample

func main() {
	template := "{{ .FirstName }} (Of the family: {{ .LastName }}) is {{ .Age }} years old\n"

	c := &config.Config{
		Format:   "gotemplate",
		Template: template,
	}
	// Single Entry
	log.Println("Printing single entry")
	person := &sample{FirstName: "Jason", LastName: "Vorhees", Age: 11}
	out, err := formatter.OutputData(person, c)
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Printing multiple entry")
	people := &sampleList{
		sample{FirstName: "Jason", LastName: "Vorhees", Age: 11},
		sample{FirstName: "Freddy", LastName: "Krueger", Age: 35},
		sample{FirstName: "Michael", LastName: "Myers", Age: 13},
	}
	out, err = formatter.OutputData(people, c)
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}
}
