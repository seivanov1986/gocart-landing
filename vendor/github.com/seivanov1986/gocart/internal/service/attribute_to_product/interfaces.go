package attribute_to_product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute_to_product"
)

type Service interface {
	Create(ctx context.Context, input AttributeToProductCreateInput) error
	Delete(ctx context.Context, ProductID, AttributeID int64) error
	List(ctx context.Context, productID int64, offset int64) (*attribute_to_product.AttributeToProductListOut, error)
}
