// Package analyzer implements the goparamorder rule as a single
// golang.org/x/tools/go/analysis.Analyzer.
package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const doc = `check parameter order against a configured list of types

This linter enforces, for every checked function signature:

- context.Context, if present, is the first parameter.
- every other parameter whose type appears in the "order" setting follows
  the relative order given there. Parameter types absent from "order",
  including every predeclared type (string, int, error, ...), are unchecked.

The "scope" setting selects which declaration kinds are checked:
"funcdecl" (function and method declarations, on by default), "funclit"
(function literals, off by default), "interface" (interface method
signatures, off by default).

Matching is by exact declared type, read from the file's own import list;
no type-checking is performed, so this analyzer only needs syntax
information.

The linter does not suggest fixes: reordering a signature would break every
call site, since Go arguments are positional.`

type goparamorder struct {
	order         []string
	checkFuncDecl bool
	checkFuncLit  bool
	checkIface    bool
}

// New creates a new goparamorder analyzer configured by s.
func New(s Settings) *analysis.Analyzer {
	enabled := s.enabledScopes()
	g := &goparamorder{
		order:         s.Order,
		checkFuncDecl: enabled[scopeFuncDecl],
		checkFuncLit:  enabled[scopeFuncLit],
		checkIface:    enabled[scopeInterface],
	}

	return &analysis.Analyzer{
		Name: "goparamorder",
		Doc:  doc,
		Run:  g.run,
	}
}

func (g *goparamorder) run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			g.inspect(pass, file, n)
			return true
		})
	}

	//nolint:nilnil // analysis.Analyzer.Run's contract: nil error means success with no shared result.
	return nil, nil
}

func (g *goparamorder) inspect(pass *analysis.Pass, file *ast.File, n ast.Node) {
	switch x := n.(type) {
	case *ast.FuncDecl:
		if g.checkFuncDecl {
			g.check(pass, file, x.Type.Params)
		}

	case *ast.FuncLit:
		if g.checkFuncLit {
			g.check(pass, file, x.Type.Params)
		}

	case *ast.InterfaceType:
		if g.checkIface {
			for _, method := range x.Methods.List {
				if funcType, ok := method.Type.(*ast.FuncType); ok {
					g.check(pass, file, funcType.Params)
				}
			}
		}
	}
}

func (g *goparamorder) check(pass *analysis.Pass, file *ast.File, params *ast.FieldList) {
	slots := flattenParams(params)
	checkContextFirst(pass, file, slots)
	checkOrder(pass, file, g.order, slots)
}
