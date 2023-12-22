package sefurl

import (
	"context"
)

type UserReadInput struct {
	ID int64
}

type UserReadRow struct {
	Login string  `db:"login" json:"login"`
	Email *string `db:"email" json:"email"`
}

func (i *repository) Read(ctx context.Context, in UserReadInput) (*UserReadRow, error) {
	row := UserReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT login, email FROM user WHERE id = ?`,
		in.ID)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
