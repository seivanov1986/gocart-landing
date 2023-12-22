package image_to_product

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in ImageToProductCreateInput) error
	DeleteImageInProduct(ctx context.Context, CategoryID, ImageID int64) error
	DeleteImagesInProduct(ctx context.Context, CategoryID int64) error
	List(ctx context.Context, in ImageToProductListInput) (*ImageToProductListOut, error)
}
