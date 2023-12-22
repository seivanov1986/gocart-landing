package attribute_to_product

import (
	"context"
)

func (i *repository) Read(ctx context.Context, ProductID, AttributeID int64) (*string, error) {
	var value string
	err := i.db.GetContext(
		ctx,
		&value,
		`SELECT value FROM attribute_to_product WHERE id_product = ? AND id_attribute`,
		ProductID, AttributeID)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
