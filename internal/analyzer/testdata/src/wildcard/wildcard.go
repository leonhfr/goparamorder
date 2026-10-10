package wildcard

import (
	"context"

	"example.com/other"
	"example.com/pkg"
)

// order (see analyzer_test.go): "example.com/pkg.B", "example.com/pkg.*",
// "example.com/other.X". B keeps its own explicit position ahead of every
// other pkg type, which the wildcard groups together before other.X.

func OK(ctx context.Context, b pkg.B, a pkg.A, x other.X) {}

func OKWildcardGroup(ctx context.Context, b pkg.B, a pkg.A, c pkg.C) {}

func BadExplicitAfterWildcard(ctx context.Context, a pkg.A, b pkg.B) {} // want `parameter of type example.com/pkg.B must come before example.com/pkg.\*, per the configured order`
