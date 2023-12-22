package image_to_category

import (
	"context"
)

type ImageToCategoryCreateInput struct {
	CategoryID int64 `db:"id_category"`
	ImageID    int64 `db:"id_image"`
}

func (u *repository) Create(ctx context.Context, in ImageToCategoryCreateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		INSERT INTO image_to_category (id_category, id_image)
		VALUES (:id_category, :id_image)
	`, in)
	return err
}
