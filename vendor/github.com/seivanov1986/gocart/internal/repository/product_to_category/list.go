package product_to_category

import (
	"context"
)

type ProductToCategoryListInput struct {
	Page int64
}

type ProductToCategoryListOut struct {
	List  []ProductToCategoryListRow
	Total int64
}

type ProductToCategoryListRow struct {
}

func (c *repository) List(ctx context.Context, in ProductToCategoryListInput) (*ProductToCategoryListOut, error) {
	imageRows := []ProductToCategoryListRow{}
	err := c.db.SelectContext(
		ctx,
		&imageRows,
		`SELECT A.id_product, A.id_category, B.name AS name_category 
			FROM product_to_category A
			LEFT JOIN category B 
			ON A.id_category = B.id
          LIMIT ?, ?`,
		in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = c.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM product_to_category`)
	if err != nil {
		return nil, err
	}

	return &ProductToCategoryListOut{
		List:  imageRows,
		Total: total,
	}, nil
}
