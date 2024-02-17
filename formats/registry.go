package formats

import (
	"sort"
)

type Creator func() Formatter

var Formats = map[string]Creator{}

// Add a new format creator
func Add(name string, creator Creator) {
	Formats[name] = creator
}

// Names returns a slice of the names of the formatters
func Names() []string {
	ret := make([]string, 0, len(Formats))
	for k := range Formats {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

// Formatter interface that defines how a thing can be formatted for output
type Formatter interface {
	Format(interface{}) ([]byte, error)
}
