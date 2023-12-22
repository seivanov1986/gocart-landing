package attribute

import (
	"context"
)

type AttributeCreateInput struct {
	Name      string  `db:"name"`
	Signature *string `db:"signature"`
}

func (u *repository) Create(ctx context.Context, in AttributeCreateInput) (*int64, error) {
	res, err := u.db.NamedExecContext(
		ctx, `
		INSERT INTO attribute (name, signature)
		VALUES (:name, :signature)
	`, in)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &lastInsertID, err
}
