package image_to_category

import (
	"context"
)

type ImageToCategoryListInput struct {
	Page       int64
	CategoryID int64
}

type ImageToCategoryListOut struct {
	List  []ImageToCategoryListRow
	Total int64
}

type ImageToCategoryListRow struct {
	ID        int64  `db:"id"`
	ImageID   int64  `db:"id_image"`
	ImagePath string `db:"path_image"`
}

func (i *repository) List(ctx context.Context, in ImageToCategoryListInput) (*ImageToCategoryListOut, error) {
	imageRows := []ImageToCategoryListRow{}
	err := i.db.SelectContext(
		ctx,
		&imageRows,
		`SELECT itc.id, itc.id_image, CONCAT(i.path, i.name) as path_image
					FROM image_to_category itc
					LEFT JOIN image i ON i.id = itc.id_image
					WHERE id_category = ? LIMIT ?, ?`,
		in.CategoryID, in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM image_to_category WHERE id_category = ?`,
		in.CategoryID)
	if err != nil {
		return nil, err
	}

	return &ImageToCategoryListOut{
		List:  imageRows,
		Total: total,
	}, nil
}
