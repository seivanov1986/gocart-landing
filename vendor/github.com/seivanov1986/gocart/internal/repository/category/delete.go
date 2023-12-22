package category

import (
	"context"
)

func (u *repository) Delete(ctx context.Context, idCategory int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM category WHERE id=?`,
		idCategory,
	)
	return err
}
