package generator

import (
	"go/types"

	"github.com/Matts966/analysisutil"
	"github.com/Matts966/genelizer/generator/hclreader"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

func generateRun(rule hclreader.Rule) func(pass *analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		if !analysisutil.PkgUsedInPass(rule.Package, pass) {
			return nil, nil
		}
		reportType(pass, rule)
		reportFunc(pass, rule)
		return nil, nil
	}
}

func reportType(pass *analysis.Pass, rule hclreader.Rule) {
	for _, rt := range rule.Types {
		t := analysisutil.TypeOf(pass, rule.Package, rt.Name)
		for _, f := range pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs {
			if !analysisutil.PkgUsedInFunc(pass, rule.Package, f) {
				continue
			}
			for _, b := range f.Blocks {
				for i := range b.Instrs {
					pos := b.Instrs[i].Pos()
					for _, rts := range rt.Shoulds {
						rtsf := analysisutil.MethodOf(t, rts)
						called, ok := analysisutil.CalledFrom(b, i, t, rtsf)
						if ok && !called {
							if rule.Message != nil {
								pass.Reportf(pos, *rule.Message)
							} else {
								pass.Reportf(pos, "should call "+rtsf.Name()+" when using "+t.String())
							}
						}
					}
				}
			}
		}
	}
}

func reportFunc(pass *analysis.Pass, rule hclreader.Rule) {
	for _, rf := range rule.Funcs {
		rfo := analysisutil.ObjectOf(pass, rule.Package, rf.Name)
		rff, ok := rfo.(*types.Func)
		if !ok {
			continue
		}
		for _, f := range pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs {
			if !analysisutil.PkgUsedInFunc(pass, rule.Package, f) {
				continue
			}
			for _, b := range f.Blocks {
				for i, instr := range b.Instrs {
					if !analysisutil.Called(instr, nil, rff) {
						continue
					}
					pos := b.Instrs[i].Pos()

					for _, rfb := range rf.Befores {
						rfbf, ok := analysisutil.ObjectOf(pass, rule.Package, rfb.Name).(*types.Func)
						if !ok {
							continue
						}
						ok, called := analysisutil.CalledFromBefore(b, i, nil, rfbf)
						if !(ok && called) {
							if rule.Message != nil {
								pass.Reportf(pos, *rule.Message)
							} else {
								pass.Reportf(pos, "should call "+rfbf.Name()+" before calling "+rf.Name)
							}
						}
					}

					for _, rfa := range rf.Afters {
						rfaf, ok := analysisutil.ObjectOf(pass, rule.Package, rfa.Name).(*types.Func)
						if !ok {
							continue
						}
						ok, called := analysisutil.CalledFromAfter(b, i, nil, rfaf)
						if !(ok && called) {
							if rule.Message != nil {
								pass.Reportf(pos, *rule.Message)
							} else {
								pass.Reportf(pos, "should call "+rfaf.Name()+" after calling "+rf.Name)
							}
						}
					}
				}
			}
		}
	}
}
