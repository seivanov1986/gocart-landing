package user

import (
	"context"
	"time"
)

type UserCreateInput struct {
	Login     string    `db:"login"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *repository) Create(ctx context.Context, in UserCreateInput) (*int64, error) {
	res, err := r.db.NamedExecContext(
		ctx, `
		INSERT INTO user (login, email, password, active, created_at)
		VALUES (:login, :email, :password, :active, :created_at)
	`, in)
	if err != nil {
		return nil, err
	}

	userLastInsertID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &userLastInsertID, err
}
