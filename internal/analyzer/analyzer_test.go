package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

var order = []string{"example.com/pkg.A", "example.com/pkg.B", "example.com/pkg.C", "basic.Local"}

func TestBasic(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, New(Settings{Order: order}), "basic")
}

func TestFuncLitScopeEnabled(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	funcLitOrder := []string{"example.com/pkg.A", "example.com/pkg.B"}
	analysistest.Run(t, testdata, New(Settings{Order: funcLitOrder, Scope: map[string]bool{scopeFuncLit: true}}), "funclit")
}

func TestFuncLitScopeDisabled(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	funcLitOrder := []string{"example.com/pkg.A", "example.com/pkg.B"}
	analysistest.Run(t, testdata, New(Settings{Order: funcLitOrder}), "funclitoff")
}

func TestInterfaceScopeEnabled(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	ifaceOrder := []string{"example.com/pkg.A", "example.com/pkg.B"}
	analysistest.Run(t, testdata, New(Settings{Order: ifaceOrder, Scope: map[string]bool{scopeInterface: true}}), "ifacescope")
}

func TestInterfaceScopeDisabled(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	ifaceOrder := []string{"example.com/pkg.A", "example.com/pkg.B"}
	analysistest.Run(t, testdata, New(Settings{Order: ifaceOrder}), "ifacescopeoff")
}

func TestFuncDeclScopeDisabled(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, New(Settings{Order: order, Scope: map[string]bool{scopeFuncDecl: false}}), "funcdecloff")
}
