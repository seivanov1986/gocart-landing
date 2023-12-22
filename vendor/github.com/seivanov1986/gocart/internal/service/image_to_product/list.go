package image_to_product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image_to_product"
)

type ImageToProductListIn struct {
	ProductID int64
	Page      int64
}

func (u *service) List(ctx context.Context, in ImageToProductListIn) (*image_to_product.ImageToProductListOut, error) {
	return u.hub.ImageToProduct().List(ctx, image_to_product.ImageToProductListInput{
		ProductID: in.ProductID,
		Page:      in.Page,
	})
}
