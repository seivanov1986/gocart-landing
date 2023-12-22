package user

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/user"
)

type UserListIn struct {
	Page int64
}

func (s *service) List(ctx context.Context, in UserListIn) (*user.UserListOut, error) {
	return s.hub.User().List(ctx, user.UserListInput{
		Page: in.Page,
	})
}
