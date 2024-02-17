# GOUT (Previously go-output-format)

[![Test](https://github.com/drewstinnett/gout/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/drewstinnett/gout/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/drewstinnett/gout/branch/main/graph/badge.svg?token=KBITDOWZLQ)](https://codecov.io/gh/drewstinnett/gout)
[![Go Report Card](https://goreportcard.com/badge/github.com/drewstinnett/gout)](https://goreportcard.com/report/github.com/drewstinnett/gout)
[![Go Reference](https://pkg.go.dev/badge/github.com/drewstinnett/gout.svg)](https://pkg.go.dev/github.com/drewstinnett/gout/v2)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

`gout` is the Go OUTput Formatter for serializing data in a standard way.

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

```go
import "github.com/drewstinnett/gout/v2"
```

### Builtin

Gout now comes with a builtin instance, similar to the way Viper does things.

Example Usage:

```go
gout.MustPrint(struct {
  FirstName string
  LastName  string
}{
  FirstName: "Bob",
  LastName:  "Ross",
})
```

Full example code [here](./_examples/builtin/main.go)

### Custom

Example Usage:

```go
// Create a new instance
w, err := gout.New()
if err != nil {
  panic(err)
}

// Use a custom writer
w.SetWriter(os.Stderr)

// Print something!
w.MustPrint(struct {
  FirstName string
  LastName  string
}{
  FirstName: "Bob",
  LastName:  "Ross",
})
// {"FirstName":"Bob","LastName":"Ross"}
```

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

### XML

Uses `encoding/xml` library. NOTE: This plugin only works with structs supported
by the base library

### Plain

This is just vanilla old Golang output, using the `%+v` format.

### GoTemplate

Use this format to parse the data in to a golang template. Useful for spitting
data out in a more arbitrary format. This uses the `text/template` package to
parse each item in the return slice. See [the example
here](_examples/gotemplate/main.go) for full details.

## Related Projects

* [Gout-Cobra](https://github.com/drewstinnett/gout-cobra) - Configure a cobra.Command to output using Gout

## Coming Soon?

### TSV (Tab Separated Values)

Intention here is to have a simple way to print out a data structure in a way
that grep and the like can parse it.
