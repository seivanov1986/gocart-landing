package page

import (
	"context"
)

type PageUpdateInput struct {
	ID           int64   `db:"id"`
	Name         string  `db:"name"`
	Content      *string `db:"content"`
	MetaID       *int64  `db:"id_meta"`
	Type         int64   `db:"type"`
	Sort         int64   `db:"sort"`
	ShortContent *string `db:"short_content"`
	ImageID      *int64  `db:"id_image"`
	CreatedAT    int64   `db:"created_at"`
	UpdatedAT    int64   `db:"updated_at"`
}

func (u *repository) Update(ctx context.Context, in PageUpdateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		UPDATE page SET name=:name, content=:content, id_meta=:id_meta,
		        type=:type, sort=:sort, short_content=:short_content, 
		        id_image=:id_image, created_at=:created_at, updated_at=:updated_at 
		WHERE id=:id
	`, in)
	return err
}
