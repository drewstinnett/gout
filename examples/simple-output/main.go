package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/drewstinnett/go-output-format/internal/formats/all"
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
	formats := formatter.GetFormats()
	formatArg := strings.Join(formats[:], "|")

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s %s", os.Args[0], formatArg)
	}
	c := &config.Config{
		Format: os.Args[1],
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
