package attribute

import (
	"context"
)

type AttributeReadInput struct {
	ID int64
}

type AttributeReadRow struct {
	Name      string  `db:"name" json:"name"`
	Sort      int64   `db:"sort" json:"sort"`
	Signature *string `db:"signature" json:"signature"`
}

func (i *repository) Read(ctx context.Context, in AttributeReadInput) (*AttributeReadRow, error) {
	row := AttributeReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT name, sort, signature FROM attribute WHERE id = ?`,
		in.ID)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
