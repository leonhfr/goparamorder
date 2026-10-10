package funcdecloff

import (
	"context"

	"example.com/pkg"
)

func BadOrder(ctx context.Context, b pkg.B, a pkg.A) {}

func BadContextNotFirst(a pkg.A, ctx context.Context) {}
