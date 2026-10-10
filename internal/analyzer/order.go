package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// paramSlot is one parameter position. A field with several names sharing
// one type, e.g. "a, b string", expands into one slot per name.
type paramSlot struct {
	pos  token.Pos
	expr ast.Expr
}

// flattenParams expands a parameter list into one slot per parameter.
func flattenParams(fields *ast.FieldList) []paramSlot {
	if fields == nil {
		return nil
	}

	var slots []paramSlot

	for _, field := range fields.List {
		if len(field.Names) == 0 {
			slots = append(slots, paramSlot{pos: field.Type.Pos(), expr: field.Type})
			continue
		}

		for _, name := range field.Names {
			slots = append(slots, paramSlot{pos: name.Pos(), expr: field.Type})
		}
	}

	return slots
}

// orderIndex resolves a type to its position among exact "import/path.Type"
// and wildcard "import/path.*" order entries.
type orderIndex struct {
	exact    map[string]int
	wildcard map[string]int
}

// newOrderIndex splits order into its exact and wildcard entries, each
// keyed for lookup by resolved type.
func newOrderIndex(order []string) orderIndex {
	idx := orderIndex{
		exact:    make(map[string]int, len(order)),
		wildcard: make(map[string]int),
	}

	for i, entry := range order {
		if pkg, ok := strings.CutSuffix(entry, ".*"); ok {
			idx.wildcard[pkg] = i
			continue
		}

		idx.exact[entry] = i
	}

	return idx
}

// lookup reports the order position of importPath.typeName. An exact entry
// takes precedence over a package wildcard.
func (idx orderIndex) lookup(importPath, typeName string) (int, bool) {
	if i, ok := idx.exact[importPath+"."+typeName]; ok {
		return i, true
	}

	i, ok := idx.wildcard[importPath]

	return i, ok
}

// checkContextFirst reports the first context.Context parameter if it does
// not sit at index 0. A later one is ignored.
func checkContextFirst(pass *analysis.Pass, file *ast.File, slots []paramSlot) {
	for i, slot := range slots {
		importPath, typeName, ok := resolveType(pass, file, slot.expr)
		if !ok || importPath != "context" || typeName != "Context" {
			continue
		}

		if i != 0 {
			pass.Reportf(slot.pos, "context.Context must be the first parameter")
		}

		return
	}
}

// checkOrder reports a parameter whose order position precedes one already
// seen earlier. Types absent from order are unchecked.
func checkOrder(pass *analysis.Pass, file *ast.File, order []string, idx orderIndex, slots []paramSlot) {
	maxSeen := -1

	for _, slot := range slots {
		importPath, typeName, ok := resolveType(pass, file, slot.expr)
		if !ok {
			continue
		}

		pos, ok := idx.lookup(importPath, typeName)
		if !ok {
			continue
		}

		if pos < maxSeen {
			pass.Reportf(slot.pos, "parameter of type %s.%s must come before %s, per the configured order", importPath, typeName, order[maxSeen])
			continue
		}

		maxSeen = pos
	}
}
