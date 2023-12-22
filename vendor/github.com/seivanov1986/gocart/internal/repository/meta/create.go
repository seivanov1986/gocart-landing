package meta

import (
	"context"
)

type MetaCreateInput struct {
	Title       *string `db:"title"`
	Keywords    *string `db:"keywords"`
	Description *string `db:"description"`
}

func (u *repository) Create(ctx context.Context, in MetaCreateInput) (*int64, error) {
	trx := u.Trx.FindTransaction(ctx)
	res, err := trx.NamedExecContext(
		ctx, `
		INSERT INTO meta (title, keywords, description)
		VALUES (:title, :keywords, :description)
	`, in)
	if err != nil {
		return nil, err
	}

	userLastInsertID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &userLastInsertID, err
}
