package attribute

import (
	"context"
)

func (r *repository) Delete(ctx context.Context, ID int64) error {
	_, err := r.db.ExecContext(
		ctx,
		`DELETE FROM attribute WHERE id = ?`,
		ID,
	)
	return err
}

func (r *repository) DeleteIn(ctx context.Context, IDs []int64) error {
	return r.db.DeleteIn(
		ctx,
		`DELETE FROM attribute WHERE id IN (?)`,
		IDs,
	)
}
