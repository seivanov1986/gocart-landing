package attribute_to_product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute_to_product"
)

func (c *service) List(ctx context.Context, productID int64, offset int64) (*attribute_to_product.AttributeToProductListOut, error) {
	return c.hub.AttributeToProduct().List(ctx, productID, offset)
}
