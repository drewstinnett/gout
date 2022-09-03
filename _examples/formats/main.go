package main

import (
	"fmt"

	gout "github.com/drewstinnett/go-output-format/v2"
	"github.com/drewstinnett/go-output-format/v2/formats/json"
	"github.com/drewstinnett/go-output-format/v2/formats/plain"
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
	formats := map[string]gout.Formatter{
		"yaml":  yaml.Formatter{},
		"plain": plain.Formatter{},
		"json":  json.Formatter{},
	}
	c, _ := gout.NewClient()
	for formatN, formatF := range formats {
		fmt.Printf("# Format: %v\n", formatN)
		c.SetFormatter(formatF)
		fmt.Println("## Person")
		c.MustPrint(person)
		fmt.Println("## People")
		c.MustPrint(people)
		fmt.Println()
	}

	// fmt.Println(string(b))
}
