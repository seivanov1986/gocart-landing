package meta

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in MetaCreateInput) (*int64, error)
	Read(ctx context.Context, in MetaReadInput) (*MetaReadRow, error)
	Update(ctx context.Context, in MetaUpdateInput) error
	Delete(ctx context.Context, idMeta int64) error
}
