package meta

import (
	"context"
)

type MetaUpdateInput struct {
	ID          int64   `db:"id"`
	Title       *string `db:"title"`
	Keywords    *string `db:"keywords"`
	Description *string `db:"description"`
}

func (u *repository) Update(ctx context.Context, in MetaUpdateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		UPDATE meta SET title=:title, keywords=:keywords, description=:description
		WHERE id=:id
	`, in)
	return err
}
