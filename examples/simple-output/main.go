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
	person := &Sample{FirstName: "Drew", LastName: "Stinnett", Age: 32}
	out, err := formatter.OutputData(person, c)
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Printing multiple entry")
	var people = &SampleList{
		Sample{FirstName: "Drew", LastName: "Stinnett", Age: 42},
		Sample{FirstName: "Lela", LastName: "Stinnett", Age: 45},
		Sample{FirstName: "James", LastName: "Stinnett", Age: 14},
	}
	out, err = formatter.OutputData(people, c)
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}

}
