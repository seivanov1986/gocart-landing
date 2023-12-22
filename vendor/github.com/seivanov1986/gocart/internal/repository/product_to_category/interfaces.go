package product_to_category

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in ProductToCategoryCreateInput) error
	DeleteIn(ctx context.Context, IDs []int64) error
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context, in ProductToCategoryListInput) (*ProductToCategoryListOut, error)
	Update(ctx context.Context, in ProductToCategoryUpdateInput) error
}
