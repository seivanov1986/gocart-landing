package page

import (
	"context"
)

type PageReadInput struct {
	ID int64
}

type PageReadRow struct {
	Name         string  `db:"name" json:"name"`
	Template     *string `db:"template" json:"template"`
	Type         int64   `db:"type" json:"type"`
	ImageID      *int64  `db:"id_image" json:"id_image"`
	ImagePath    *string `db:"path_image" json:"path_image"`
	ShortContent *string `db:"short_content" json:"short_content"`
	Content      *string `db:"content" json:"content"`
	Sort         int64   `db:"sort" json:"sort"`
	SefUrl       string  `db:"sefurl" json:"sefurl"`
	Title        *string `db:"title" json:"title"`
	Keywords     *string `db:"keywords" json:"keywords"`
	Description  *string `db:"description" json:"description"`
	MetaID       *int64  `db:"id_meta" json:"id_meta"`
}

func (i *repository) Read(ctx context.Context, in PageReadInput) (*PageReadRow, error) {
	row := PageReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT p.name, p.content, p.type, 
       			p.sort, p.short_content, p.id_image, p.id_meta,
       			m.title, m.keywords, m.description,
       			s.name as sefurl, s.template,
       			CONCAT(i.path, i.name) as path_image
			FROM page p
			LEFT JOIN meta m ON m.id = p.id_meta
			LEFT JOIN sefurl s ON s.id_object = p.id AND s.type = 1
			LEFT JOIN image i ON i.id = p.id_image
			WHERE p.id = ?`,
		in.ID)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
