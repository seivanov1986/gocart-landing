package page

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/page"
)

type PageReadIn struct {
	ID int64
}

func (u *service) Read(ctx context.Context, in PageReadIn) (*page.PageReadRow, error) {
	return u.hub.Page().Read(ctx, page.PageReadInput{
		ID: in.ID,
	})
}
