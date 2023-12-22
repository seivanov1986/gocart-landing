package page

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/page"
)

type Service interface {
	Create(ctx context.Context, in PageCreateIn) (*int64, error)
	Read(ctx context.Context, in PageReadIn) (*page.PageReadRow, error)
	Update(ctx context.Context, in PageUpdateInput) error
	Delete(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in PageListIn) (*page.PageListOut, error)
}
