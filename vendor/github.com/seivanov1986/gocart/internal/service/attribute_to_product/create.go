package attribute_to_product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute_to_product"
)

type AttributeToProductCreateInput struct {
	ProductID   int64  `db:"id_product"`
	AttributeID int64  `db:"id_attribute"`
	Value       string `db:"value"`
}

func (c *service) Create(ctx context.Context, input AttributeToProductCreateInput) error {
	return c.hub.AttributeToProduct().Create(ctx, attribute_to_product.AttributeToProductCreateInput{
		ProductID:   input.ProductID,
		AttributeID: input.AttributeID,
		Value:       input.Value,
	})
}
