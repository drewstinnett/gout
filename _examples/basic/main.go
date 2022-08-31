package main

import (
	"os"

	writer "github.com/drewstinnett/go-output-format/v2"
	"github.com/drewstinnett/go-output-format/v2/formats/json"
)

func main() {
	w, err := writer.NewClient()
	if err != nil {
		panic(err)
	}
	// By Default, the YAML format is use, Let's change it to json though
	w.SetFormatter(json.Formatter{})

	// By Default, print to stdout. Let's change it to stderr though
	w.SetWriter(os.Stderr)

	// Print it on out!
	w.Print(struct {
		FirstName string
		LastName  string
	}{
		FirstName: "Bob",
		LastName:  "Ross",
	})
	// {"FirstName":"Bob","LastName":"Ross"}
}
