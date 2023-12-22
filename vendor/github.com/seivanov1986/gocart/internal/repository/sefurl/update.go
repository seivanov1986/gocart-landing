package sefurl

import (
	"context"
)

type SefUrlUpdateInput struct {
	Url       string  `db:"url"`
	Path      string  `db:"path"`
	Name      string  `db:"name"`
	Type      int64   `db:"type"`
	ObjectID  int64   `db:"id_object"`
	Template  *string `db:"template"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

func (u *repository) Update(ctx context.Context, in SefUrlUpdateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		UPDATE sefurl SET url=:url, path=:path, name=:name, template=:template, 
		                created_at=:created_at, updated_at=:updated_at 
		WHERE id_object=:id_object AND type=:type
	`, in)
	return err
}
