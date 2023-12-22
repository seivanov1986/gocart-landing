package attribute_to_product

import (
	"context"
)

type AttributeToProductCreateInput struct {
	ProductID   int64  `db:"id_product"`
	AttributeID int64  `db:"id_attribute"`
	Value       string `db:"value"`
}

func (u *repository) Create(ctx context.Context, in AttributeToProductCreateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		INSERT INTO attribute_to_product (id_product, id_attribute, value)
		VALUES (:id_product, :id_attribute, :value)
	`, in)
	return err
}
