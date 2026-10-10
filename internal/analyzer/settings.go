package analyzer

import "maps"

// Settings configures the analyzer.
type Settings struct {
	// Order lists type paths ("import/path.TypeName") in the order their
	// parameters must appear. Do not list context.Context; see the doc.
	Order []string `json:"order"`

	// Scope selects which declaration kinds to check: "funcdecl" (on by
	// default), "funclit", "interface" (both off). See defaultScopes.
	Scope map[string]bool `json:"scope"`
}

const (
	scopeFuncDecl  = "funcdecl"
	scopeFuncLit   = "funclit"
	scopeInterface = "interface"
)

var defaultScopes = map[string]bool{
	scopeFuncDecl:  true,
	scopeFuncLit:   false,
	scopeInterface: false,
}

// enabledScopes resolves the settings into the effective set of scopes to
// check, applying s.Scope as overrides on top of defaultScopes.
func (s Settings) enabledScopes() map[string]bool {
	enabled := make(map[string]bool, len(defaultScopes))
	maps.Copy(enabled, defaultScopes)
	maps.Copy(enabled, s.Scope)

	return enabled
}
