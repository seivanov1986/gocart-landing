package category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/category"
)

type Service interface {
	Create(ctx context.Context, in CategoryCreateIn) (*int64, error)
	Read(ctx context.Context, in CategoryReadIn) (*category.CategoryReadRow, error)
	Update(ctx context.Context, in CategoryUpdateInput) error
	Delete(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in CategoryListIn) (*category.CategoryListOut, error)
}
