package attribute

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type AttributeCreateIn struct {
	Name      string  `db:"name"`
	Signature *string `db:"signature"`
}

func (u *service) Create(ctx context.Context, in AttributeCreateIn) (*int64, error) {
	return u.hub.Attribute().Create(ctx, attribute.AttributeCreateInput{
		Name:      in.Name,
		Signature: in.Signature,
	})
}
