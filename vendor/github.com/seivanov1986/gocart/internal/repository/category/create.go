package category

import (
	"context"
)

type CategoryCreateInput struct {
	Name      string  `db:"name"`
	ParentID  *int64  `db:"id_parent"`
	Content   *string `db:"content"`
	ImageID   *int64  `db:"id_image"`
	MetaID    *int64  `db:"id_meta"`
	Sort      int64   `db:"sort"`
	Price     *string `db:"price"`
	Disabled  bool    `db:"disabled"`
	CreatedAT int64   `db:"created_at"`
	UpdatedAT int64   `db:"updated_at"`
}

func (u *repository) Create(ctx context.Context, in CategoryCreateInput) (*int64, error) {
	trx := u.Trx.FindTransaction(ctx)
	res, err := trx.NamedExecContext(
		ctx, `
		INSERT INTO category (name, id_parent, content, id_image, id_meta, sort, price, disabled, created_at, updated_at)
		VALUES (:name, :id_parent, :content, :id_image, :id_meta, :sort, :price, :disabled, :created_at, :updated_at)
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
