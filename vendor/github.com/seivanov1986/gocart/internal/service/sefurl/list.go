package sefurl

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type SefUrlListIn struct {
	Page int64
}

func (u *service) List(ctx context.Context, in SefUrlListIn) (*sefurl.SefUrlListOut, error) {
	return u.hub.SefUrl().List(ctx, sefurl.SefUrlListInput{
		Page: in.Page,
	})
}
