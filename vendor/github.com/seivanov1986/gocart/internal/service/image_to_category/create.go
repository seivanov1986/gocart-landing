package image_to_category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image_to_category"
)

type ImageToCategoryCreateIn struct {
	CategoryID int64 `db:"id_category"`
	ImageID    int64 `db:"id_image"`
}

func (u *service) Create(ctx context.Context, in ImageToCategoryCreateIn) error {
	return u.hub.ImageToCategory().Create(ctx, image_to_category.ImageToCategoryCreateInput{
		CategoryID: in.CategoryID,
		ImageID:    in.ImageID,
	})
}
