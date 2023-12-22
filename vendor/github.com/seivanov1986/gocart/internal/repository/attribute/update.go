package attribute

import (
	"context"
)

type AttributeUpdateInput struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	Signature *string `db:"signature"`
}

func (u *repository) Update(ctx context.Context, in AttributeUpdateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		UPDATE attribute SET name=:name, signature=:signature WHERE id=:id
	`, in)
	return err
}
