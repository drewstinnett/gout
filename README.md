# go-output-format

[![Test](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/drewstinnett/go-output-format/branch/main/graph/badge.svg?token=KBITDOWZLQ)](https://codecov.io/gh/drewstinnett/go-output-format)
[![Go Report Card](https://goreportcard.com/badge/github.com/drewstinnett/go-output-format)](https://goreportcard.com/report/github.com/drewstinnett/go-output-format)
[![Go Reference](https://pkg.go.dev/badge/github.com/drewstinnett/go-output-format.svg)](https://pkg.go.dev/github.com/drewstinnett/go-output-format)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

*NOTE*: V2 is a breaking compatibility change from V1. Going forward, only V2 will
be developed and supported.

Helper utility to output data structures in to standardized formats, much like
what is built in to [vault](https://www.vaultproject.io/),
[az](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) and
[kubectl](https://kubernetes.io/docs/tasks/tools/)

I really like how these apps provide for flexible output, but wanted a way to
do it without needing to re-write or copy it for every new tool.

Need to parse some output with `jq`?  JSON is your format. Want to put
it out in an easy to read yet still standardized format?  YAML is for you!

This tool is intended to provide all that in a single reusable package.

## Usage

Import with:

```go
import "github.com/drewstinnett/go-output-format/v2/writer"
```

Example Usage:

```go
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
```

See [_examples](_examples/) for more example usage

## Built-in Formatters

### YAML

Uses the standard `gopkg.in/yaml.v3` library.

### JSON

Uses the standard `encoding/json` library.

### Plain

This is just vanilla old Golang output, using the `%+v` format.

### GoTemplate

Use this format to parse the data in to a golang template. Useful for spitting
data out in a more arbitrary format. This uses the `text/template` package to
parse each item in the return slice. See [the example
here](examples/templated-output/main.go) for full details.

## Coming Soon/TODO

### TSV (Tab Separated Values)

Intention here is to have a simple way to print out a data structure in a way
that grep and the like can parse it.

### Cobra Integrations

To make it as easy as possible to integrate in CLI apps, I'd like to add some
bindings for automatically creating the writer from
[cobra](https://github.com/spf13/cobra) arguments
