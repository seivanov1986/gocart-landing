package user

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/user"
)

type UserUpdateInput struct {
	ID       int64  `db:"id"`
	Password string `db:"password"`
}

func (s *service) Update(ctx context.Context, in UserUpdateInput) error {
	return s.hub.User().Update(ctx, user.UserUpdateInput{
		ID:       in.ID,
		Password: "",
	})
}
