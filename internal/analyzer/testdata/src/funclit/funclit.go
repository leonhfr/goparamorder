package funclit

import (
	"context"

	"example.com/pkg"
)

func Caller() {
	f := func(a pkg.A, ctx context.Context) {} // want `context.Context must be the first parameter`
	f(pkg.A{}, context.Background())

	g := func(ctx context.Context, b pkg.B, a pkg.A) {} // want `parameter of type example.com/pkg.A must come before example.com/pkg.B, per the configured order`
	g(context.Background(), pkg.B{}, pkg.A{})
}
