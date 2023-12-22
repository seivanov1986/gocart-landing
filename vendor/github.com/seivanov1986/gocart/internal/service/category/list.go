package category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/category"
)

type CategoryListIn struct {
	Page     int64
	ParentID int64
}

func (u *service) List(ctx context.Context, in CategoryListIn) (*category.CategoryListOut, error) {
	return u.hub.Category().List(ctx, category.CategoryListInput{
		Page:     in.Page,
		ParentID: in.ParentID,
	})
}
