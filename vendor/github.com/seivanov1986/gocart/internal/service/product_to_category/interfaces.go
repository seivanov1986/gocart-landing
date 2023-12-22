package product_to_category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/product_to_category"
)

type Service interface {
	Create(ctx context.Context, input ProductToCategoryCreateInput) error
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context, in ProductToCategoryListInput) (*product_to_category.ProductToCategoryListOut, error)
}
