package product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/product"
)

type ProductListIn struct {
	Page int64
}

func (u *service) List(ctx context.Context, in ProductListIn) (*product.ProductListOut, error) {
	return u.hub.Product().List(ctx, product.ProductListInput{
		Page: in.Page,
	})
}
