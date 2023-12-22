package attribute_to_product

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in AttributeToProductCreateInput) error
	Read(ctx context.Context, ProductID, AttributeID int64) (*string, error)
	Update(ctx context.Context, in UserUpdateInput) error
	Delete(ctx context.Context, ProductID, AttributeID int64) error
	List(ctx context.Context, productID int64, offset int64) (*AttributeToProductListOut, error)
}
