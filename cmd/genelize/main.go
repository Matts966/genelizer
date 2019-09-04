package main

import (
	"github.com/Matts966/genelizer/generator"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() { multichecker.Main(generator.Generate()...) }
