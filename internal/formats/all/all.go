package all

import (
	// Plugins need to register themselves
	_ "github.com/drewstinnett/go-output-format/internal/formats/gotemplate"
	_ "github.com/drewstinnett/go-output-format/internal/formats/yaml"
)
