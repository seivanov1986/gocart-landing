package image_to_product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image_to_product"
)

type Service interface {
	Create(ctx context.Context, input ImageToProductCreateInput) error
	Delete(ctx context.Context, CategoryID, ImageID int64) error
	List(ctx context.Context, in ImageToProductListIn) (*image_to_product.ImageToProductListOut, error)
}
