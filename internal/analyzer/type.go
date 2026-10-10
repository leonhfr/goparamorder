package analyzer

import (
	"go/ast"
	"path"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

// predeclared holds every predeclared type identifier. Their parameter
// position is never checked.
var predeclared = map[string]bool{
	"bool": true, "string": true, "error": true, "any": true,
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true, "uintptr": true,
	"float32": true, "float64": true,
	"complex64": true, "complex128": true,
	"byte": true, "rune": true,
}

// resolveType reports the import path and type name a parameter type
// names, stripping one pointer level; ok is false if unresolvable.
func resolveType(pass *analysis.Pass, file *ast.File, expr ast.Expr) (importPath, typeName string, ok bool) {
	if star, isStar := expr.(*ast.StarExpr); isStar {
		expr = star.X
	}

	switch t := expr.(type) {
	case *ast.Ident:
		if predeclared[t.Name] {
			return "", "", false
		}

		return pass.Pkg.Path(), t.Name, true

	case *ast.SelectorExpr:
		pkgIdent, isIdent := t.X.(*ast.Ident)
		if !isIdent {
			return "", "", false
		}

		importPath, found := resolveImport(file, pkgIdent.Name)
		if !found {
			return "", "", false
		}

		return importPath, t.Sel.Name, true

	default:
		return "", "", false
	}
}

// resolveImport finds name's import path in file, by alias, or by its
// default package name (last path segment) when unaliased.
func resolveImport(file *ast.File, name string) (string, bool) {
	for _, imp := range file.Imports {
		importPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}

		alias := path.Base(importPath)
		if imp.Name != nil {
			alias = imp.Name.Name
		}

		if alias == name {
			return importPath, true
		}
	}

	return "", false
}
