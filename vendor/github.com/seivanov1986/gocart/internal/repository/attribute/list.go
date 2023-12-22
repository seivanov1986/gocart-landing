package attribute

import (
	"context"
)

type AttributeListInput struct {
	Page int64
}

type AttributeListOut struct {
	List  []AttributeListRow
	Total int64
}

type AttributeListRow struct {
	ID        int64   `db:"id" json:"id"`
	Name      string  `db:"name" json:"name"`
	Sort      int64   `db:"sort" json:"sort"`
	Signature *string `db:"signature" json:"signature"`
}

func (i *repository) List(ctx context.Context, in AttributeListInput) (*AttributeListOut, error) {
	pageRows := []AttributeListRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, name, sort, signature FROM attribute LIMIT ?, ?`,
		in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM attribute`)
	if err != nil {
		return nil, err
	}

	return &AttributeListOut{
		List:  pageRows,
		Total: total,
	}, nil
}
