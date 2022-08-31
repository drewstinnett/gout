package formatter

import (
	"github.com/drewstinnett/go-output-format/pkg/config"
)

type Client struct {
	Config *config.Config
}

// NewClient returns a new formatter client, given a config and io writer
func NewClient(c *config.Config) (*Client, error) {
	return nil, nil
}
