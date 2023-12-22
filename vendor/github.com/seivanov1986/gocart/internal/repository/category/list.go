package category

import (
	"context"
)

type CategoryListInput struct {
	Page     int64
	ParentID int64
}

type CategoryListOut struct {
	List  []CategoryListRow
	Total int64
}

type CategoryListRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

const (
	limit = 8
)

func (i *repository) List(ctx context.Context, in CategoryListInput) (*CategoryListOut, error) {
	pageRows := []CategoryListRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, name FROM category WHERE id_parent = ? LIMIT ?, ?`,
		in.ParentID, in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM category WHERE id_parent = ?`,
		in.ParentID,
	)
	if err != nil {
		return nil, err
	}

	return &CategoryListOut{
		List:  pageRows,
		Total: total,
	}, nil
}
