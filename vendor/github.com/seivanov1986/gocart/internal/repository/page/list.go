package page

import (
	"context"
)

type PageListInput struct {
	Page int64
}

type PageListOut struct {
	List  []PageListRow
	Total int64
}

type PageListRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Type int64  `db:"type" json:"type"`
}

const (
	limit = 8
)

func (i *repository) List(ctx context.Context, in PageListInput) (*PageListOut, error) {
	pageRows := []PageListRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, name, type FROM page LIMIT ?, ?`,
		in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM page`)
	if err != nil {
		return nil, err
	}

	return &PageListOut{
		List:  pageRows,
		Total: total,
	}, nil
}
