package product

import (
	"context"
)

type ProductUpdateInput struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	Content   *string `db:"content"`
	ImageID   *int64  `db:"id_image"`
	MetaID    *int64  `db:"id_meta"`
	Sort      int64   `db:"sort"`
	Price     *string `db:"price"`
	Disabled  bool    `db:"disabled"`
	CreatedAT int64   `db:"created_at"`
	UpdatedAT int64   `db:"updated_at"`
}

func (u *repository) Update(ctx context.Context, in ProductUpdateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		UPDATE product SET name=:name, content=:content, id_meta=:id_meta,
		        id_image=:id_image, sort=:sort, price=:price, disabled=:disabled, 
		         created_at=:created_at, updated_at=:updated_at 
		WHERE id=:id
	`, in)
	return err
}
