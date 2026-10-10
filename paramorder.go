// Package goparamorder provides a linter that enforces a configured
// parameter order in function declarations. See the internal/analyzer
// package for the rule implementation.
package goparamorder

import (
	"golang.org/x/tools/go/analysis"

	"github.com/leonhfr/goparamorder/internal/analyzer"
)

// Settings configures the analyzer.
type Settings = analyzer.Settings

// New creates a new goparamorder analyzer configured by s.
func New(s Settings) *analysis.Analyzer {
	return analyzer.New(s)
}
