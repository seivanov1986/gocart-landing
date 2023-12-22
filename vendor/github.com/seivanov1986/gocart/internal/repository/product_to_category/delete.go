package product_to_category

import (
	"context"
)

func (u *repository) DeleteIn(ctx context.Context, IDs []int64) error {
	return u.db.DeleteIn(
		ctx,
		`DELETE FROM product_to_category WHERE id IN (?)`,
		IDs,
	)
}

func (u *repository) Delete(ctx context.Context, ID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM product_to_category WHERE id = ?`,
		ID,
	)
	return err
}
