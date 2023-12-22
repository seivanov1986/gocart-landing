package product

import (
	"context"
)

func (u *repository) Delete(ctx context.Context, IDs []int64) error {
	return u.db.DeleteIn(
		ctx,
		`DELETE FROM product WHERE id IN (?)`,
		IDs,
	)
}
