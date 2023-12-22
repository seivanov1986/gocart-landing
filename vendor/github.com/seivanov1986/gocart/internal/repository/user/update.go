package user

import (
	"context"
)

type UserUpdateInput struct {
	ID       int64  `db:"id"`
	Password string `db:"password"`
}

func (r *repository) Update(ctx context.Context, in UserUpdateInput) error {
	_, err := r.db.NamedExecContext(
		ctx, `
		UPDATE user SET password=:password WHERE id=:id
	`, in)
	return err
}
