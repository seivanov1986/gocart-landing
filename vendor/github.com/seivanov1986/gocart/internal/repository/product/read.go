package product

import (
	"context"
)

type ProductReadInput struct {
	ID int64
}

type ProductReadRow struct {
	Name        string  `db:"name" json:"name"`
	Content     *string `db:"content" json:"content"`
	MetaID      *int64  `db:"id_meta" json:"id_meta"`
	Sort        int64   `db:"sort" json:"sort"`
	Price       *int64  `db:"price" json:"price"`
	ImageID     *int64  `db:"id_image" json:"id_image"`
	ImagePath   *string `db:"path_image" json:"path_image"`
	Disabled    bool    `db:"disabled" json:"disabled"`
	Template    *string `db:"template" json:"template"`
	SefUrl      string  `db:"sefurl" json:"sefurl"`
	Title       *string `db:"title" json:"title"`
	Keywords    *string `db:"keywords" json:"keywords"`
	Description *string `db:"description" json:"description"`
}

func (i *repository) Read(ctx context.Context, in ProductReadInput) (*ProductReadRow, error) {
	row := ProductReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT p.name, p.content, p.id_meta, 
       			p.sort, p.price, p.id_image, p.disabled,  
       			m.title, m.keywords, m.description,
       			s.name as sefurl, s.template,
       			CONCAT(i.path, i.name) as path_image
			FROM product p
			LEFT JOIN meta m ON m.id = p.id_meta
			LEFT JOIN sefurl s ON s.id_object = p.id AND s.type = 3
			LEFT JOIN image i ON i.id = p.id_image
			WHERE p.id = ?`,
		in.ID)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
