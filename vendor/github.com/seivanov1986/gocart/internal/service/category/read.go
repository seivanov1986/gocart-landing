package category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/category"
)

type CategoryReadIn struct {
	ID int64
}

func (u *service) Read(ctx context.Context, in CategoryReadIn) (*category.CategoryReadRow, error) {
	return u.hub.Category().Read(ctx, in.ID)
}
