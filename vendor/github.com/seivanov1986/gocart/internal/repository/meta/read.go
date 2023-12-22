package meta

import (
	"context"
)

type MetaReadInput struct {
	ID int64
}

type MetaReadRow struct {
	Title       *string `db:"title" json:""`
	Keywords    *string `db:"keywords"`
	Description *string `db:"description"`
}

func (i *repository) Read(ctx context.Context, in MetaReadInput) (*MetaReadRow, error) {
	row := MetaReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT name, content, id_meta, type, 
       			sort, short_content, id_image, created_at, updated_at 
			FROM page WHERE id = ?`,
		in.ID)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
