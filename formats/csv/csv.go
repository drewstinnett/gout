package csv

import (
	"github.com/jszwec/csvutil"
)

type Formatter struct{}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	return csvutil.Marshal(v)
}
