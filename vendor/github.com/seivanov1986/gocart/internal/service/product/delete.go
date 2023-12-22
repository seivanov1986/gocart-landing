package product

import (
	"context"
)

func (u *service) Delete(ctx context.Context, IDs []int64) error {
	return u.hub.Product().Delete(ctx, IDs)
}
