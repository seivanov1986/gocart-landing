package page

import (
	"context"
)

type PageCreateInput struct {
	Name         string  `db:"name"`
	Content      *string `db:"content"`
	MetaID       *int64  `db:"id_meta"`
	Type         int64   `db:"type"`
	Sort         int64   `db:"sort"`
	ShortContent *string `db:"short_content"`
	ImageID      *int64  `db:"id_image"`
	CreatedAT    int64   `db:"created_at"`
	UpdatedAT    int64   `db:"updated_at"`
}

func (u *repository) Create(ctx context.Context, in PageCreateInput) (*int64, error) {
	trx := u.Trx.FindTransaction(ctx)
	res, err := trx.NamedExecContext(
		ctx, `
		INSERT INTO page (name, content, id_meta, type, sort, short_content, id_image, created_at, updated_at)
		VALUES (:name, :content, :id_meta, :type, :sort, :short_content, :id_image, :created_at, :updated_at)
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
