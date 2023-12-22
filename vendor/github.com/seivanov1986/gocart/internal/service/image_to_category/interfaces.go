package image_to_category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/image_to_category"
)

type Service interface {
	Create(ctx context.Context, in ImageToCategoryCreateIn) error
	Delete(ctx context.Context, categoryID int64) error
	List(ctx context.Context, in ImageToCategoryListIn) (*image_to_category.ImageToCategoryListOut, error)
}
