[![Test](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/drewstinnett/go-output-format/branch/main/graph/badge.svg?token=KBITDOWZLQ)](https://codecov.io/gh/drewstinnett/go-output-format)
[![Go Report Card](https://goreportcard.com/badge/github.com/drewstinnett/go-output-format)](https://goreportcard.com/report/github.com/drewstinnett/go-output-format)
[![Go Reference](https://pkg.go.dev/badge/github.com/drewstinnett/go-output-format.svg)](https://pkg.go.dev/github.com/drewstinnett/go-output-format)

# go-output-format

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
import "github.com/drewstinnett/go-output-format/formatter"
```

Example Usage:
```go
package main

import (
    "fmt"

    "github.com/drewstinnett/go-output-format/formatter"
)

func main() {
    c := &formatter.Config{
        Format: "yaml",
    }

    type sample struct {
        FirstName string
        LastName  string
        Age       int
    }

    person := &sample{FirstName: "Jason", LastName: "Vorhees", Age: 11}
    out, err := formatter.OutputData(person, c)

    fmt.Println(string(out))
}
```

See [examples](examples/) for more example usage

## Formatter details

### YAML

Uses the standard `gopkg.in/yaml.v2` library.

### JSON

Uses the standard `encoding/json` library.

### Plain

This is just vanilla old Golang output, using the `%+v` format.

### GoTemplate

Use this format to parse the data in to a golang template. Useful for spitting
data out in a more arbitrary format. This uses the `text/template` package to
parse each item in the return slice. See [the example
here](examples/templated-output/main.go) for full details.

### TSV (Tab Separated Values)

Here be dragons, this is one I wrote. Intention here is to have a simple way to
print out a data structure in a way that grep and the like can parse it.
