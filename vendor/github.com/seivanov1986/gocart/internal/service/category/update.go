package category

import (
	"context"
	"fmt"
	"time"

	"github.com/seivanov1986/gocart/internal/repository/category"
	"github.com/seivanov1986/gocart/internal/repository/meta"
	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type CategoryUpdateInput struct {
	ID           int64   `db:"id"`
	Name         string  `db:"name"`
	ParentID     *int64  `db:"id_parent"`
	Content      *string `db:"content"`
	MetaID       *int64  `db:"id_meta"`
	Type         int64   `db:"type"`
	Sort         int64   `db:"sort"`
	ShortContent *string `db:"short_content"`
	ImageID      *int64  `db:"id_image"`
	SefURL       string  `db:"sefurl"`
	Template     *string `db:"template"`
	Title        *string `db:"title"`
	Keywords     *string `db:"keywords"`
	Description  *string `db:"description"`
}

func (u *service) Update(ctx context.Context, in CategoryUpdateInput) error {
	updatedAt := time.Now()

	var parentID int64 = 0
	if in.ParentID != nil {
		parentID = *in.ParentID
	}

	return u.TrManager.MakeTransaction(ctx, func(ctx context.Context) error {
		// check parentID

		row, err := u.hub.Category().Read(ctx, in.ID)
		if err != nil {
			return fmt.Errorf("page read", err.Error())
		}

		var metaID = row.MetaID
		if row.MetaID != nil {
			err := u.hub.Meta().Update(ctx, meta.MetaUpdateInput{
				ID:          *row.MetaID,
				Title:       in.Title,
				Keywords:    in.Keywords,
				Description: in.Description,
			})
			if err != nil {
				return err
			}
		} else {
			metaID, err = u.hub.Meta().Create(ctx, meta.MetaCreateInput{
				Title:       in.Title,
				Keywords:    in.Keywords,
				Description: in.Description,
			})
			if err != nil {
				return err
			}
		}

		// TODO: check if not sefurl -> create
		// TODO: transaction manager

		err = u.hub.SefUrl().Update(ctx, sefurl.SefUrlUpdateInput{
			Url:      "/" + in.SefURL,
			Path:     "/",
			Name:     in.SefURL,
			Type:     categoryType,
			ObjectID: in.ID,
			Template: in.Template,
		})
		if err != nil {
			return err
		}

		err = u.hub.Category().Update(ctx, category.CategoryUpdateInput{
			ParentID:  &parentID,
			ID:        in.ID,
			Name:      in.Name,
			Content:   in.Content,
			Sort:      in.Sort,
			ImageID:   in.ImageID,
			UpdatedAT: updatedAt.Unix(),
			MetaID:    metaID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
