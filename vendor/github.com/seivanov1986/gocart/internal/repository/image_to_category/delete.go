package image_to_category

import (
	"context"
)

func (u *repository) DeleteImageInCategory(ctx context.Context, CategoryID, ImageID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM image_to_category WHERE id_category=? and id_image=?`,
		CategoryID, ImageID,
	)
	return err
}

func (u *repository) DeleteImagesInCategory(ctx context.Context, CategoryID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM image_to_category WHERE id_category=?`,
		CategoryID,
	)
	return err
}
