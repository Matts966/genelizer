package generator

import (
	"go/token"
	"log"
	"path/filepath"

	"github.com/Matts966/analysisutil"
	"github.com/Matts966/genelizer/generator/hclreader"
	"github.com/gobuffalo/packr"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

// Generate generates a slice of pointers to analysis.Analyzer by rules defined in config.
func Generate() []*analysis.Analyzer {
	box := packr.NewBox(filepath.Join(hclreader.Parent, hclreader.ConfigDir))

	var analyzers []*analysis.Analyzer

	for _, conf := range box.List() {
		s, err := box.FindString(conf)
		if err != nil {
			log.Fatal(err)
		}
		config, err := hclreader.Read(s)
		if err != nil {
			log.Fatal(err)
		}
		// if config == nil {
		// 	log.Fatal(fmt.Errorf("empty config"))
		// }
		for _, rule := range config {
			var analyzer = &analysis.Analyzer{
				Name: rule.Name,
				RunDespiteErrors: true,
				Requires: []*analysis.Analyzer{
					buildssa.Analyzer,
				},
			}
			if rule.Doc != nil {
				analyzer.Doc = *rule.Doc
			}

			if rule.Package != nil {
				analyzer.Run = generateRun(*rule.Package)
			}

			analyzers = append(analyzers, analyzer)
		}
	}

	return analyzers
}

func generateRun(packageName string) func(pass *analysis.Pass) (interface{}, error) {

	log.Println(packageName)

	return func(pass *analysis.Pass) (interface{}, error) {
		if analysisutil.PkgUsedInPass(packageName, pass) {
			pass.Reportf(token.NoPos, packageName+" package is used in pass.")
		}
		return nil, nil
	}
}
