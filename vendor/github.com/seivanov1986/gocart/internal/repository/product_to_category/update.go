package product_to_category

import (
	"context"
)

type ProductToCategoryUpdateInput struct {
	ID           int64  `db:"id"`
	ProductID    int64  `db:"id_product"`
	CategoryID   int64  `db:"id_category"`
	MainCategory *int64 `db:"main_category"`
}

func (u *repository) Update(ctx context.Context, in ProductToCategoryUpdateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		UPDATE product_to_category SET id_product=:id_product, 
			id_category=:id_category, main_category=:main_category WHERE id=:id
	`, in)
	return err
}
