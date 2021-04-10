[![Test](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/drewstinnett/go-output-format/branch/main/graph/badge.svg?token=KBITDOWZLQ)](https://codecov.io/gh/drewstinnett/go-output-format)

# go-output-format

Generically format output for CLI apps, inspired by Hashicorp Vaults CLI
options.

When using CLI based tools, it's often useful to output the data in different
formats. Need to parse some output with `jq`?  JSON is your format. Want to put
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

### TSV (Tab Separated Values)

Here be dragons, this is one I wrote. Intention here is to have a simple way to
print out a data structure in a way that grep and the like can parse it. It's a
little junky right now, but give it time 😁

