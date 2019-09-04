package hclreader

import (
	"github.com/hashicorp/hcl"
)

// ConfigDir is a path to directory that contains config files.
const (
	Parent    = ".."
	ConfigDir = "config"
)

// Read inputs the path of file and read it as hcl.
func Read(s string) (Config, error) {
	var c Config
	err := hcl.Decode(c, s)
	if err != nil {
		return nil, err
	}
	return c, nil
}
