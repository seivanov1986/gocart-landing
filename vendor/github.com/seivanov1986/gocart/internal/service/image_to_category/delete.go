package image_to_category

import (
	"context"
)

func (c *service) Delete(ctx context.Context, categoryID int64) error {
	return c.hub.ImageToCategory().DeleteImagesInCategory(ctx, categoryID)
}
