package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/drewstinnett/go-output-format/formatter"
)

type Sample struct {
	FirstName string
	LastName  string
	Age       int
}

type SampleList []Sample

func main() {
	formats := formatter.GetFormats()
	formatArg := strings.Join(formats[:], "|")

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s %s", os.Args[0], formatArg)
	}
	c := &formatter.FormatterConfig{
		Format: os.Args[1],
	}
	// Single Entry
	log.Println("Printing single entry")
	person := &Sample{FirstName: "Jason", LastName: "Vorhees", Age: 11}
	out, err := formatter.OutputData(person, c)
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Printing multiple entry")
	var people = &SampleList{
		Sample{FirstName: "Jason", LastName: "Vorhees", Age: 11},
		Sample{FirstName: "Freddy", LastName: "Krueger", Age: 35},
		Sample{FirstName: "Michael", LastName: "Myers", Age: 13},
	}
	out, err = formatter.OutputData(people, c)
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}

}
