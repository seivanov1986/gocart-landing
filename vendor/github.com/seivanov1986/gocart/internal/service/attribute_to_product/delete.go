package attribute_to_product

import (
	"context"
)

func (c *service) Delete(ctx context.Context, ProductID, AttributeID int64) error {
	return c.hub.AttributeToProduct().Delete(ctx, ProductID, AttributeID)
}
