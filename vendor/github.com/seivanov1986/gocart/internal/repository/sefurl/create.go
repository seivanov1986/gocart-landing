package sefurl

import (
	"context"
)

type SefUrlCreateInput struct {
	Url       string  `db:"url"`
	Path      string  `db:"path"`
	Name      string  `db:"name"`
	Type      int64   `db:"type"`
	ObjectID  int64   `db:"id_object"`
	Template  *string `db:"template"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

func (u *repository) Create(ctx context.Context, in SefUrlCreateInput) (*int64, error) {
	trx := u.Trx.FindTransaction(ctx)
	res, err := trx.NamedExecContext(
		ctx, `
		INSERT INTO sefurl (url, path, name, type, id_object, template, created_at, updated_at)
		VALUES (:url, :path, :name, :type, :id_object, :template, :created_at, :updated_at)
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
