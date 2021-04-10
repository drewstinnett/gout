[![Test](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/drewstinnett/go-output-format/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/drewstinnett/go-output-format/branch/main/graph/badge.svg?token=KBITDOWZLQ)](https://codecov.io/gh/drewstinnett/go-output-format)

# go-output-format

Generically format output for CLI apps, inspired by Hashicorp Vaults CLI
options.

When using CLI based tools, it's often useful to output the data in different
formats. I tend to use the following:

* YAML: Great for human readability, but still in a standardized format
* JSON: Super ugly to read, but awesome when combined with tooling like `jq`
* TSV: Great for just simple quick and easy grep/awk/sed/etc

This tool is intended to provide all that in a single reusable package

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
