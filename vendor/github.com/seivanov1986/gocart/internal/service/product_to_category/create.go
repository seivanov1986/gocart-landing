package product_to_category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/product_to_category"
)

type ProductToCategoryCreateInput struct {
	ProductID    int64  `db:"id_product"`
	CategoryID   int64  `db:"id_category"`
	MainCategory *int64 `db:"main_category"`
}

func (c *service) Create(ctx context.Context, input ProductToCategoryCreateInput) error {
	return c.hub.ProductToCategory().Create(ctx, product_to_category.ProductToCategoryCreateInput{
		ProductID:    input.ProductID,
		CategoryID:   input.CategoryID,
		MainCategory: input.MainCategory,
	})
}
