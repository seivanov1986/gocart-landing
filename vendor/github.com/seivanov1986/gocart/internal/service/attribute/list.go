package attribute

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type AttributeListIn struct {
	Page int64
}

func (u *service) List(ctx context.Context, in AttributeListIn) (*attribute.AttributeListOut, error) {
	return u.hub.Attribute().List(ctx, attribute.AttributeListInput{
		Page: in.Page,
	})
}
