package hclreader

import (
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"

	"golang.org/x/xerrors"
)

// ConfigDir is a path to directory that contains config files.
const (
	Parent    = ".."
	ConfigDir = "config"
)

// Read inputs the path of file and read it as hcl.
func Read(b []byte) (*Config, error) {
	parser := hclparse.NewParser()
	f, parseDiags := parser.ParseHCL(b, "config")
	if parseDiags.HasErrors() {
		return nil, xerrors.Errorf("parsing hcl failed: %s", parseDiags.Error())
	}

	var c Config
	decodeDiags := gohcl.DecodeBody(f.Body, nil, &c)
	if decodeDiags.HasErrors() {
		return nil, xerrors.Errorf("decoding hcl failed: %s", decodeDiags.Error())
	}

	return &c, nil
}
