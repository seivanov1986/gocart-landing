package product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/product"
)

type ProductReadIn struct {
	ID int64
}

func (u *service) Read(ctx context.Context, in ProductReadIn) (*product.ProductReadRow, error) {
	return u.hub.Product().Read(ctx, product.ProductReadInput{
		ID: in.ID,
	})
}
