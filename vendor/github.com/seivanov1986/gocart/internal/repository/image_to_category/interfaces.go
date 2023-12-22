package image_to_category

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in ImageToCategoryCreateInput) error
	DeleteImageInCategory(ctx context.Context, CategoryID, ImageID int64) error
	DeleteImagesInCategory(ctx context.Context, CategoryID int64) error
	List(ctx context.Context, in ImageToCategoryListInput) (*ImageToCategoryListOut, error)
}
