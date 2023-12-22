package user

import (
	"context"
)

func (r *repository) PasswordByLogin(ctx context.Context, login string) (*string, error) {
	var password string
	err := r.db.GetContext(
		ctx,
		&password,
		`SELECT password FROM user WHERE login = ?`,
		login)
	if err != nil {
		return nil, err
	}

	return &password, nil
}
