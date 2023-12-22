package attribute

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in AttributeCreateInput) (*int64, error)
	Read(ctx context.Context, in AttributeReadInput) (*AttributeReadRow, error)
	Update(ctx context.Context, in AttributeUpdateInput) error
	Delete(ctx context.Context, ID int64) error
	DeleteIn(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in AttributeListInput) (*AttributeListOut, error)
}
