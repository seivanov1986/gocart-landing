package image_to_category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image_to_category"
)

type ImageToCategoryListIn struct {
	CategoryID int64
	Page       int64
}

func (u *service) List(ctx context.Context, in ImageToCategoryListIn) (*image_to_category.ImageToCategoryListOut, error) {
	return u.hub.ImageToCategory().List(ctx, image_to_category.ImageToCategoryListInput{
		Page:       in.Page,
		CategoryID: in.CategoryID,
	})
}
