package user

import (
	"context"
)

type UserReadRow struct {
	Login string  `db:"login" json:"login"`
	Email *string `db:"email" json:"email"`
}

func (r *repository) Read(ctx context.Context, ID int64) (*UserReadRow, error) {
	row := UserReadRow{}
	err := r.db.GetContext(
		ctx,
		&row,
		`SELECT login, email FROM user WHERE id = ?`,
		ID)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
