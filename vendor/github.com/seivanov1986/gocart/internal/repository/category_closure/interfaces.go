package category_closure

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, idObject, idParent int64) error
}
