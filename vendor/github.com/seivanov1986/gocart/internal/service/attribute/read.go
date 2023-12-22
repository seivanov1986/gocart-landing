package attribute

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type AttributeReadIn struct {
	ID int64
}

func (u *service) Read(ctx context.Context, in AttributeReadIn) (*attribute.AttributeReadRow, error) {
	return u.hub.Attribute().Read(ctx, attribute.AttributeReadInput{
		ID: in.ID,
	})
}
