package product

import (
	"context"
)

type ProductCreateInput struct {
	Name      string  `db:"name"`
	Content   *string `db:"content"`
	ImageID   *int64  `db:"id_image"`
	MetaID    *int64  `db:"id_meta"`
	Sort      int64   `db:"sort"`
	Price     *string `db:"price"`
	Disabled  bool    `db:"disabled"`
	CreatedAT int64   `db:"created_at"`
	UpdatedAT int64   `db:"updated_at"`
}

func (u *repository) Create(ctx context.Context, in ProductCreateInput) (*int64, error) {
	trx := u.Trx.FindTransaction(ctx)
	res, err := trx.NamedExecContext(
		ctx, `
		INSERT INTO product (name, content, id_image, id_meta, sort, price, disabled, created_at, updated_at)
		VALUES (:name, :content, :id_image, :id_meta, :sort, :price, :disabled, :created_at, :updated_at)
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
