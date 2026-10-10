// Package goparamorder registers the goparamorder linter as a
// golangci-lint module plugin.
//
// Build a custom golangci-lint binary that includes it by listing this
// module in .custom-gcl.yml and running `golangci-lint custom`; then enable
// the linter under linters.settings.custom in .golangci.yml.
package goparamorder

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/leonhfr/goparamorder/internal/analyzer"
)

func init() {
	register.Plugin("goparamorder", newPlugin)
}

type plugin struct {
	settings Settings
}

func newPlugin(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &plugin{settings: s}, nil
}

// BuildAnalyzers returns the single analyzer this plugin provides.
func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.New(p.settings)}, nil
}

// GetLoadMode requires only syntax: matching is by exact declared type,
// never by type-checked interface satisfaction.
func (p *plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
