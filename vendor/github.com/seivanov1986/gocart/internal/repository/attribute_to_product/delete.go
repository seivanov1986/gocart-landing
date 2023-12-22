package attribute_to_product

import (
	"context"
)

func (u *repository) Delete(ctx context.Context, ProductID, AttributeID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM attribute_to_product WHERE id_product=? AND id_attribute=?`,
		ProductID, AttributeID,
	)
	return err
}
