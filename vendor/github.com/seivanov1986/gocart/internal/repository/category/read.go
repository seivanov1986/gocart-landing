package category

import (
	"context"
)

type CategoryReadRow struct {
	Name        string  `db:"name" json:"name"`
	ParentID    *int64  `db:"id_parent" json:"id_parent"`
	Content     *string `db:"content" json:"content"`
	ImageID     *int64  `db:"id_image" json:"id_image"`
	ImagePath   *string `db:"path_image" json:"path_image"`
	MetaID      *int64  `db:"id_meta" json:"id_meta"`
	Sort        int64   `db:"sort" json:"sort"`
	Price       *int64  `db:"price" json:"price"`
	Disabled    bool    `db:"disabled" json:"disabled"`
	Template    *string `db:"template" json:"template"`
	SefUrl      string  `db:"sefurl" json:"sefurl"`
	Title       *string `db:"title" json:"title"`
	Keywords    *string `db:"keywords" json:"keywords"`
	Description *string `db:"description" json:"description"`
}

func (i *repository) Read(ctx context.Context, categoryID int64) (*CategoryReadRow, error) {
	row := CategoryReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT c.name, c.id_parent, c.content, c.id_image,
       			c.id_meta, c.sort, c.price, c.disabled,  
       			m.title, m.keywords, m.description,
       			s.name as sefurl, s.template,
       			CONCAT(i.path, i.name) as path_image
			FROM category c
			LEFT JOIN meta m ON m.id = c.id_meta
			LEFT JOIN sefurl s ON s.id_object = c.id AND s.type = 2
			LEFT JOIN image i ON i.id = c.id_image
			WHERE c.id = ?`,
		categoryID)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
