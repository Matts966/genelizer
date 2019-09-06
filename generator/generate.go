package generator

import (
	"fmt"
	"log"

	"github.com/Matts966/genelizer/generator/hclreader"
	"github.com/gobuffalo/packr/v2"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

const confDir = "../config"

// Generate generates a slice of pointers to analysis.Analyzer by rules defined in config.
func Generate() []*analysis.Analyzer {
	box := packr.New("Config", confDir)

	var analyzers []*analysis.Analyzer

	for _, conf := range box.List() {
		b, err := box.Find(conf)
		if err != nil {
			log.Fatal(err)
		}
		config, err := hclreader.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		if config == nil {
			log.Fatal(fmt.Errorf("empty config"))
		}

		for _, rule := range config.Rules {
			var analyzer = &analysis.Analyzer{
				Name:             rule.Name,
				RunDespiteErrors: true,
				Doc:              rule.Doc,
				Requires: []*analysis.Analyzer{
					buildssa.Analyzer,
				},
			}
			analyzer.Run = generateRun(rule)
			analyzers = append(analyzers, analyzer)
		}
	}

	return analyzers
}
