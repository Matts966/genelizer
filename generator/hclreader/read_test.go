package hclreader_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Matts966/genelizer/generator/hclreader"
	"golang.org/x/xerrors"
)

const confDir = "../../config"

func TestRead(t *testing.T) {
	confDir := confDir

	err := filepath.Walk(confDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip itself
		if path == confDir {
			return nil
		}

		if info.IsDir() {
			return xerrors.New("directory is not allowed in config directory")
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return xerrors.Errorf("reading "+path+" failed, err: %+v", err)
		}

		_, err = hclreader.Read(b)
		if err != nil {
			return xerrors.Errorf("decoding "+path+" failed, err: %+v", err)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
