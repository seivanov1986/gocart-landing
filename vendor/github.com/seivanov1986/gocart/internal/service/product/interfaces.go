package product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/product"
)

type Service interface {
	Create(ctx context.Context, in ProductCreateIn) (*int64, error)
	Read(ctx context.Context, in ProductReadIn) (*product.ProductReadRow, error)
	Update(ctx context.Context, in ProductUpdateInput) error
	Delete(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in ProductListIn) (*product.ProductListOut, error)
}
