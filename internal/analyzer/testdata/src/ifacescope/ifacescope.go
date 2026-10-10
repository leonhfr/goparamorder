package ifacescope

import (
	"context"

	"example.com/pkg"
)

type Service interface {
	Do(a pkg.A, ctx context.Context)             // want `context.Context must be the first parameter`
	Other(ctx context.Context, b pkg.B, a pkg.A) // want `parameter of type example.com/pkg.A must come before example.com/pkg.B, per the configured order`
}
