package attribute

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type Service interface {
	Create(ctx context.Context, in AttributeCreateIn) (*int64, error)
	Read(ctx context.Context, in AttributeReadIn) (*attribute.AttributeReadRow, error)
	Update(ctx context.Context, in AttributeUpdateInput) error
	Delete(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in AttributeListIn) (*attribute.AttributeListOut, error)
}
