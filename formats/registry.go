/*
Package formats provides some base extractions for format plugins
*/
package formats

import (
	"sort"
)

// Creator is a generator for new formatter instances
type Creator func() Formatter

// Formats defines a map of labels to Creators
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
