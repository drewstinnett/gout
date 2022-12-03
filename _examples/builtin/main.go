package main

import (
	"github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/formats/json"
)

func main() {
	gout.SetFormatter(json.Formatter{})
	gout.MustPrint(struct {
		FirstName string
		LastName  string
	}{
		FirstName: "Bob",
		LastName:  "Ross",
	})
	// {"FirstName":"Bob","LastName":"Ross"}
}
