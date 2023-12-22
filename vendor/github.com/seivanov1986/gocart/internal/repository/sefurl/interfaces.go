package sefurl

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in SefUrlCreateInput) (*int64, error)
	Read(ctx context.Context, in UserReadInput) (*UserReadRow, error)
	Update(ctx context.Context, in SefUrlUpdateInput) error
	DeleteByObjectType(ctx context.Context, ObjectID, Type int64) error
	List(ctx context.Context, in SefUrlListInput) (*SefUrlListOut, error)
}
