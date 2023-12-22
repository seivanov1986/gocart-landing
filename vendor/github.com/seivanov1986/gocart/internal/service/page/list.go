package page

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/page"
)

type PageListIn struct {
	Page int64
}

func (u *service) List(ctx context.Context, in PageListIn) (*page.PageListOut, error) {
	return u.hub.Page().List(ctx, page.PageListInput{
		Page: in.Page,
	})
}
