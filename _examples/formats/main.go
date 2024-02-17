package main

import (
	"fmt"

	gout "github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/formats"
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
	g := gout.New()
	for formatN, formatG := range formats.Formats {
		fmt.Printf("# Format: %v\n", formatN)
		g.SetFormatter(formatG())
		// CSV Formatter won't work on a single object, has to be iterable
		if formatN != "csv" {
			fmt.Println("## Person")
			g.MustPrint(person)
		}
		fmt.Println("## People")
		g.MustPrint(people)
		fmt.Println()
	}
}
