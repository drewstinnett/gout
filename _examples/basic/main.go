package main

import (
	"fmt"
	"os"

	gout "github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/formats"
	_ "github.com/drewstinnett/gout/v2/formats/builtin"
)

func main() {
	w := gout.New()
	fmt.Printf("Active Formatters: %v\n", formats.Names())
	// By Default, the YAML format is use, Let's change it to json though
	// w.SetFormatter(json.Formatter{})

	// By Default, print to stdout. Let's change it to stderr though
	w.SetWriter(os.Stderr)

	// Print it on out!
	w.MustPrint(struct {
		FirstName string
		LastName  string
	}{
		FirstName: "Bob",
		LastName:  "Ross",
	})
	// {"FirstName":"Bob","LastName":"Ross"}
}
