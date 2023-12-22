package attribute_to_product

import (
	"context"
)

type UserUpdateInput struct {
	ID       int64  `db:"id"`
	Password string `db:"password"`
}

func (u *repository) Update(ctx context.Context, in UserUpdateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		UPDATE attribute_to_product SET password=:password WHERE id=:id
	`, in)
	return err
}
