package ifacescopeoff

import (
	"context"

	"example.com/pkg"
)

type Service interface {
	Do(a pkg.A, ctx context.Context)
	Other(ctx context.Context, b pkg.B, a pkg.A)
}
