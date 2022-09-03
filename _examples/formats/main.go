package main

import (
	"fmt"

	"github.com/drewstinnett/go-output-format/v2/gout"
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
	c, _ := gout.New()
	for formatN, formatF := range gout.BuiltInFormatters {
		if formatN != "gotemplate" {
			fmt.Printf("# Format: %v\n", formatN)
			c.SetFormatter(formatF)
			// CSV Formatter won't work on a single object, has to be iterable
			if formatN != "csv" {
				fmt.Println("## Person")
				c.MustPrint(person)
			}
			fmt.Println("## People")
			c.MustPrint(people)
			fmt.Println()
		}
	}

	// fmt.Println(string(b))
}
