package product_to_category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/product_to_category"
)

type ProductToCategoryListInput struct {
	Page int64
}

func (u *service) List(ctx context.Context, in ProductToCategoryListInput) (*product_to_category.ProductToCategoryListOut, error) {
	return u.hub.ProductToCategory().List(ctx, product_to_category.ProductToCategoryListInput{
		Page: in.Page,
	})
}
