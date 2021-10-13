package all

import (
	// Plugins need to register themselves
	_ "github.com/drewstinnett/go-output-format/internal/formats/gotemplate"
	_ "github.com/drewstinnett/go-output-format/internal/formats/json"
	_ "github.com/drewstinnett/go-output-format/internal/formats/plain"
	_ "github.com/drewstinnett/go-output-format/internal/formats/tsv"
	_ "github.com/drewstinnett/go-output-format/internal/formats/yaml"
)
