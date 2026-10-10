package funclitoff

import (
	"context"

	"example.com/pkg"
)

func Caller() {
	f := func(a pkg.A, ctx context.Context) {}
	f(pkg.A{}, context.Background())

	g := func(ctx context.Context, b pkg.B, a pkg.A) {}
	g(context.Background(), pkg.B{}, pkg.A{})
}
