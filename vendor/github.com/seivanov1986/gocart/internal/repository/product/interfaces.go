package product

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in ProductCreateInput) (*int64, error)
	Read(ctx context.Context, in ProductReadInput) (*ProductReadRow, error)
	Update(ctx context.Context, in ProductUpdateInput) error
	Delete(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in ProductListInput) (*ProductListOut, error)
}
