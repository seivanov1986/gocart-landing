package image_to_product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image_to_product"
)

type ImageToProductCreateInput struct {
	ProductID int64 `db:"id_product"`
	ImageID   int64 `db:"id_image"`
}

func (c *service) Create(ctx context.Context, input ImageToProductCreateInput) error {
	return c.hub.ImageToProduct().Create(ctx, image_to_product.ImageToProductCreateInput{
		ProductID: input.ProductID,
		ImageID:   input.ImageID,
	})
}
