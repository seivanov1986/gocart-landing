package image_to_product

import (
	"context"
)

func (u *repository) DeleteImageInProduct(ctx context.Context, ProductID, ImageID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM image_to_product WHERE id_product=? and id_image=?`,
		ProductID, ImageID,
	)
	return err
}

func (u *repository) DeleteImagesInProduct(ctx context.Context, ProductID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM image_to_product WHERE id_product=?`,
		ProductID,
	)
	return err
}
