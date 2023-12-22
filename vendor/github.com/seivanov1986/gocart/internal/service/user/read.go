package user

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/user"
)

func (s *service) Read(ctx context.Context, ID int64) (*user.UserReadRow, error) {
	return s.hub.User().Read(ctx, ID)
}
