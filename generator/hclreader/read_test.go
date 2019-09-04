package hclreader_test

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"testing"
	"github.com/Matts966/genelizer/generator/hclreader"
)

func TestRead(t *testing.T) {
	confDir := filepath.Join(hclreader.Parent, hclreader.Parent, hclreader.ConfigDir)

	err := filepath.Walk(confDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
		}

		// Skip itself
		if path == confDir {
			return nil
		}

        if info.IsDir() {
            return fmt.Errorf("directory is not allowed in config directory")
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading "+path+" failed, err: %+v", err)
		}

		c, err := hclreader.Read(string(b))
		if err != nil {
			return fmt.Errorf("decoding "+path+" failed, err: %+v", err)
		}

		err = validate(t, c)
		if err != nil {
			return fmt.Errorf("validating "+path+" failed, err: %+v", err)
		}

        return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func validate(t *testing.T, c hclreader.Config) error {
	if c == nil {
		return fmt.Errorf("no configuration")
	}

	t.Log(c)

	for _, rule := range c {

		t.Log(rule.Name)

		if rule.Name == "" {
			
			return fmt.Errorf("rule.name is required")
		}
	}
	return nil
}
