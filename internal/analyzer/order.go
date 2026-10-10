package analyzer

import (
	"go/ast"
	"go/token"

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

// checkContextFirst reports the first context.Context parameter if it does
// not sit at index 0. A later one is ignored.
func checkContextFirst(pass *analysis.Pass, file *ast.File, slots []paramSlot) {
	for i, slot := range slots {
		typePath, ok := resolveType(pass, file, slot.expr)
		if !ok || typePath != "context.Context" {
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
func checkOrder(pass *analysis.Pass, file *ast.File, order []string, slots []paramSlot) {
	index := make(map[string]int, len(order))
	for i, typePath := range order {
		index[typePath] = i
	}

	maxSeen := -1

	for _, slot := range slots {
		typePath, ok := resolveType(pass, file, slot.expr)
		if !ok {
			continue
		}

		pos, ok := index[typePath]
		if !ok {
			continue
		}

		if pos < maxSeen {
			pass.Reportf(slot.pos, "parameter of type %s must come before %s, per the configured order", typePath, order[maxSeen])
			continue
		}

		maxSeen = pos
	}
}
