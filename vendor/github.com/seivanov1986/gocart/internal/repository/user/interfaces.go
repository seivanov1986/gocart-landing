package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in UserCreateInput) (*int64, error)
	Read(ctx context.Context, ID int64) (*UserReadRow, error)
	Update(ctx context.Context, in UserUpdateInput) error
	Delete(ctx context.Context, ID int64) error
	DeleteIn(ctx context.Context, IDs []int64) error
	List(ctx context.Context, in UserListInput) (*UserListOut, error)
	PasswordByLogin(ctx context.Context, login string) (*string, error)
}
