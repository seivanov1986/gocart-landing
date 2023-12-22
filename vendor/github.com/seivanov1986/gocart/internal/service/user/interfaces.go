package user

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/user"
)

type Service interface {
	Create(ctx context.Context, in UserCreateIn) (*int64, error)
	Read(ctx context.Context, ID int64) (*user.UserReadRow, error)
	Update(ctx context.Context, in UserUpdateInput) error
	Delete(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in UserListIn) (*user.UserListOut, error)
}
