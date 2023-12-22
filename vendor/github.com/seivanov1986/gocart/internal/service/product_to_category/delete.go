package product_to_category

import (
	"context"
)

func (c *service) Delete(ctx context.Context, ID int64) error {
	return c.hub.ProductToCategory().Delete(ctx, ID)
}
