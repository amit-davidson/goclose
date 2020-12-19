package statJanitor

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "statJanitor",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

const (
	Doc = "statJanitor finds objects that didn't call Close()"

	ioPath = "io"
	closerInterface = "Closer"
)

type runner struct {
	pass      *analysis.Pass
	closeTyp  *types.Interface
	closeMthd *types.Func
	skipFile  map[*ast.File]bool
}

func run(pass *analysis.Pass) (interface{}, error) {
	runner := runner{}
	return runner.run(pass)
}

func (r *runner) isCloserType(t types.Type) bool {
	return types.Implements(t, r.closeTyp) || types.Identical(t, r.closeTyp)
}

func (r *runner) isCloserFunc(c *types.Func) bool {
	return c.Pkg() == r.closeMthd.Pkg() && c.Name() == r.closeMthd.Name()
}

func (r *runner) getClosingType(pkgs []*ssa.Package) types.Type {
	for _, pkg := range pkgs {
		if pkg.Pkg.Name() == ioPath {
			for memName, mem := range pkg.Members {
				if memName == closerInterface {
					return mem.Type()
				}
			}
		}
	}
	return nil
}

func (r *runner) run(pass *analysis.Pass) (interface{}, error) {
	r.pass = pass
	ssaAnalysis := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	funcs := ssaAnalysis.SrcFuncs
	t := r.getClosingType(ssaAnalysis.Pkg.Prog.AllPackages())
	if t == nil {
		// skip checking
		return nil, nil
	}
	r.closeTyp, _ = t.(*types.Named).Underlying().(*types.Interface)
	r.closeMthd = r.closeTyp.Method(0)
	for _, f := range funcs {
		for _, b := range f.Blocks {
			for _, i := range b.Instrs {
				v, ok := i.(ssa.Value)
				if !ok {
					continue
				}
				if !v.Pos().IsValid() {
					continue
				}
				t := v.Type()
				if r.isCloserType(t) {
					if !r.isCloserExists(v.Referrers()) {
						pass.Reportf(v.Pos(), "response body must be closed")
					}
				}
			}
		}
	}

	return nil, nil
}

func (r *runner) isClosureCalled(is *[]ssa.Instruction) bool {
	for _, i := range *is {
		if _, ok := i.(*ssa.Call); ok {
			return true
		}
	}
	return false
}

func (r *runner) isCloserExists(is *[]ssa.Instruction) bool {
	for _, i := range *is {
		switch v := i.(type) {
		case *ssa.Call:
			methodCall := v.Call.Method
			if methodCall != nil && r.isCloserFunc(methodCall) {
				return true // Reached close function
			} else {
				function := v.Call.StaticCallee()
				var ioCloser *ssa.Parameter
				for _, p := range function.Params {
					t := p.Type()
					if r.isCloserType(t) {
						ioCloser = p
					}
				}
				if ioCloser != nil {
					return r.isCloserExists(ioCloser.Referrers())
				} else {
					fmt.Println("Can't happen func")
				}
			}
		case *ssa.MakeClosure:
			// We assume that if the closing function is passed as an argument to a function, then it'll be used inside
			return r.isClosureCalled(v.Referrers())
		case *ssa.ChangeInterface:
			return r.isCloserExists(v.Referrers())
		}
	}
	return false
}
