package wildcardreverse

import (
	"context"

	"example.com/other"
	"example.com/pkg"
)

// order (see analyzer_test.go): "example.com/pkg.*", "example.com/pkg.B",
// "example.com/other.X". The wildcard is listed before B, but B still
// keeps its own explicit position: other pkg types must precede it.

func OK(ctx context.Context, a pkg.A, b pkg.B, x other.X) {}

func BadWildcardAfterExplicit(ctx context.Context, b pkg.B, a pkg.A) {} // want `parameter of type example.com/pkg.A must come before example.com/pkg.B, per the configured order`
