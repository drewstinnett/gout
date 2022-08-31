package main

import (
	"log"

	writer "github.com/drewstinnett/go-output-format/v2"
	"github.com/drewstinnett/go-output-format/v2/formats/json"
	"github.com/drewstinnett/go-output-format/v2/formats/plain"
	"github.com/drewstinnett/go-output-format/v2/formats/tsv"
	"github.com/drewstinnett/go-output-format/v2/formats/yaml"
)

type sample struct {
	FirstName string
	LastName  string
	Age       int
}

type sampleList []sample

func main() {
	person := &sample{FirstName: "Jason", LastName: "Vorhees", Age: 11}
	people := &sampleList{
		sample{FirstName: "Jason", LastName: "Vorhees", Age: 11},
		sample{FirstName: "Freddy", LastName: "Krueger", Age: 35},
		sample{FirstName: "Michael", LastName: "Myers", Age: 13},
	}
	formats := map[string]writer.Formatter{
		"yaml":  yaml.Formatter{},
		"tsv":   tsv.Formatter{},
		"plain": plain.Formatter{},
		"json":  json.Formatter{},
	}
	c, _ := writer.NewClient()
	for formatN, formatF := range formats {
		log.Printf("Format: %v\n", formatN)
		c.SetFormatter(formatF)
		c.MustPrint(person)
		c.MustPrint(people)
	}
	// fmt.Println(string(b))
}