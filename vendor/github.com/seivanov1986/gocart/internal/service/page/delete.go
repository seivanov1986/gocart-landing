package page

import (
	"context"
)

func (u *service) Delete(ctx context.Context, IDs []int64) error {
	return u.hub.Page().Delete(ctx, IDs)
}
