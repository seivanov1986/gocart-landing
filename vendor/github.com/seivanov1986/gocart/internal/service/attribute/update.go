package attribute

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type AttributeUpdateInput struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	Signature *string `db:"signature"`
}

func (u *service) Update(ctx context.Context, in AttributeUpdateInput) error {
	return u.hub.Attribute().Update(ctx, attribute.AttributeUpdateInput{
		ID:        in.ID,
		Name:      in.Name,
		Signature: in.Signature,
	})
}
