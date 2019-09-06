package generator_test

import (
	"testing"

	"github.com/Matts966/genelizer/generator"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	for _, a := range generator.Generate() {
		analysistest.Run(t, testdata, a, "a")
	}
}
