package generator

import (
	"go/types"
	"strings"

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
					var rtsfs []*types.Func
					for _, rts := range rt.Shoulds {
						rtsfs = append(rtsfs, analysisutil.MethodOf(t, rts))
					}
					called, ok := analysisutil.CalledFrom(b, i, t, rtsfs...)
					if ok && !called {
						if rule.Message != nil {
							pass.Reportf(pos, *rule.Message)
						} else {
							pass.Reportf(pos, "should call "+strings.Join(rt.Shoulds, " or ")+" when using "+t.String())
						}
					}
				}
			}
		}
	}
}

func reportFunc(pass *analysis.Pass, rule hclreader.Rule) {
	for _, rf := range rule.Funcs {
		var recvT types.Type
		var rff *types.Func
		if rf.Receiver != nil {
			recvT = analysisutil.TypeOf(pass, rule.Package, *rf.Receiver)
			rff = analysisutil.MethodOf(recvT, rf.Name)
		} else {
			// TODO(Matts966): More precisely get object.
			rfo := analysisutil.ObjectOf(pass, rule.Package, rf.Name)
			var ok bool
			rff, ok = rfo.(*types.Func)
			if !ok {
				continue
			}
		}
		for _, f := range pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs {
			if !analysisutil.PkgUsedInFunc(pass, rule.Package, f) {
				continue
			}
			for _, b := range f.Blocks {
				for i, instr := range b.Instrs {
					recv := analysisutil.ReturnReceiverIfCalled(instr, rff)
					if recvT != nil && recv == nil {
						continue
					}

					pos := b.Instrs[i].Pos()

					for _, rfb := range rf.Befores {
						var rfbf *types.Func
						if recvT != nil {
							rfbf = analysisutil.MethodOf(recvT, rfb.Name)
						} else {
							// TODO(Matts966): More precisely get object.
							rfo := analysisutil.ObjectOf(pass, rule.Package, rfb.Name)
							var ok bool
							rfbf, ok = rfo.(*types.Func)
							if !ok {
								continue
							}
						}

						ok, called := analysisutil.CalledFromBefore(b, i, recv, rfbf)
						if !(ok && called) {
							if rule.Message != nil {
								pass.Reportf(pos, *rule.Message)
							} else {
								pass.Reportf(pos, "should call "+rfbf.Name()+" before calling "+rf.Name)
							}
						}

						for _, rfbr := range rfb.Returns {
							o := analysisutil.ObjectOf(pass, rule.Package, rfbr)
							if !analysisutil.CalledBeforeAndEqualTo(b, recv, rfbf, o) {
								pass.Reportf(pos, rfbf.Name()+" should be "+o.Name()+" when calling "+rf.Name)
							}
						}
					}

					for _, rfa := range rf.Afters {
						rfaf, ok := analysisutil.ObjectOf(pass, rule.Package, rfa.Name).(*types.Func)
						if !ok {
							continue
						}
						ok, called := analysisutil.CalledFromAfter(b, i, recv, rfaf)
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
