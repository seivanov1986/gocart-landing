package attribute_to_product

import (
	"context"
)

type AttributeToProductListOut struct {
	List  []AttributeToProductListRow
	Total int64
}

type AttributeToProductListRow struct {
	Id          int64   `db:"id"`
	IdProduct   int64   `db:"id_product"`
	IdAttribute int64   `db:"id_attribute"`
	Name        string  `db:"name"`
	Value       *string `db:"value"`
}

func (c *repository) List(ctx context.Context, productID int64, offset int64) (*AttributeToProductListOut, error) {
	pageRows := []AttributeToProductListRow{}
	var attributeToProductOut *AttributeToProductListOut

	err := c.db.SelectContext(ctx,
		&pageRows,
		`SELECT atp.id, atp.id_product, atp.id_attribute, a.name, atp.value 
		FROM attribute_to_product atp
		JOIN attribute a ON atp.id_attribute = a.id
		WHERE atp.id_product = ? LIMIT ?, ?`,
		productID,
		offset,
		limit,
	)

	var total int64
	err = c.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM attribute_to_product WHERE id_product = ?`,
		productID)
	if err != nil {
		return nil, err
	}

	return &AttributeToProductListOut{
		List:  pageRows,
		Total: total,
	}, nil

	return attributeToProductOut, nil
}
