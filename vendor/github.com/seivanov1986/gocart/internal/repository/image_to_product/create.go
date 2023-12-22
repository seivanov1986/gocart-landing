package image_to_product

import (
	"context"
)

type ImageToProductCreateInput struct {
	ProductID int64 `db:"id_product"`
	ImageID   int64 `db:"id_image"`
}

func (u *repository) Create(ctx context.Context, in ImageToProductCreateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		INSERT INTO image_to_product (id_product, id_image)
		VALUES (:id_product, :id_image)
	`, in)
	return err
}
