/*
Package builtin is a helper to  register all the built in formatters
*/
package builtin

import (
	_ "github.com/drewstinnett/gout/v2/formats/gotemplate"
	_ "github.com/drewstinnett/gout/v2/formats/json"
	_ "github.com/drewstinnett/gout/v2/formats/yaml"
)
