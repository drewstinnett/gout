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

### Basic

Import with:

```go
import "github.com/drewstinnett/go-output-format/v2/gout"
```

Example Usage:

```go
import (
 "os"
 "github.com/drewstinnett/go-output-format/v2/gout"
 "github.com/drewstinnett/go-output-format/v2/formats/json"
)

func main() {
 w, err := gout.New()
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

### Cobra Integration

To simplify using this in new projects, you can use the `NewWithCobraCmd`
method. Example:

```go
// By default, look for a field called 'format'
w, err := NewWithCobraCmd(cmd, nil)
```

```go
// Or pass a configuration object with what the field is called
w, err := NewWithCobraCmd(cmd, &gout.CobraCmdConfig{
        FormatField: "my-special-name-field",
})
```

By default, the gout will use os.Stdout as the default writer.

See [_examples](_examples/) for more example usage

## Built-in Formatters

### YAML

Uses the standard `gopkg.in/yaml.v3` library.

### JSON

Uses the standard `encoding/json` library.

### TOML

Uses `github.com/pelletier/go-toml/v2` library

### CSV

Uses `github.com/jszwec/csvutil` library. NOTE: You must pass an iterable
interface in when using this format. It won't do a single struct.

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
