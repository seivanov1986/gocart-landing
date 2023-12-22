package product_to_category

import (
	"context"
)

type ProductToCategoryCreateInput struct {
	ProductID    int64  `db:"id_product"`
	CategoryID   int64  `db:"id_category"`
	MainCategory *int64 `db:"main_category"`
}

func (u *repository) Create(ctx context.Context, in ProductToCategoryCreateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		INSERT INTO product_to_category (id_product, id_category, main_category)
		VALUES (:id_product, :id_category, :main_category)
	`, in)
	return err
}
