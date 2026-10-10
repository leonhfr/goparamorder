package basic

import (
	"context"

	"example.com/pkg"
)

type Local struct{}

type Unlisted struct{}

func OK(ctx context.Context, a pkg.A, b pkg.B, c pkg.C) {}

func OKWithPrimitives(ctx context.Context, name string, a pkg.A, n int, b pkg.B) {}

func OKPointer(ctx context.Context, a *pkg.A, b pkg.B) {}

func OKUnlisted(ctx context.Context, a pkg.A, u Unlisted, b pkg.B) {}

func OKNoContext(a pkg.A, b pkg.B) {}

func OKMultiContext(ctx, ctx2 context.Context, a pkg.A) {}

func OKGroupedNames(ctx context.Context, a, b pkg.A) {}

func BadContextNotFirst(a pkg.A, ctx context.Context) {} // want `context.Context must be the first parameter`

func BadOrder(ctx context.Context, b pkg.B, a pkg.A) {} // want `parameter of type example.com/pkg.A must come before example.com/pkg.B, per the configured order`

func BadLocalOrder(ctx context.Context, l Local, c pkg.C) {} // want `parameter of type example.com/pkg.C must come before basic.Local, per the configured order`

func BadUnnamedOrder(context.Context, pkg.B, pkg.A) {} // want `parameter of type example.com/pkg.A must come before example.com/pkg.B, per the configured order`
