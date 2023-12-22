package image_to_product

import (
	"context"
)

func (c *service) Delete(ctx context.Context, CategoryID, ImageID int64) error {
	return c.hub.ImageToProduct().DeleteImageInProduct(ctx, CategoryID, ImageID)
}
