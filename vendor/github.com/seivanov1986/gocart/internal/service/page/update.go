package page

import (
	"context"
	"fmt"
	"time"

	meta2 "github.com/seivanov1986/gocart/internal/repository/meta"
	"github.com/seivanov1986/gocart/internal/repository/page"
	sefurl2 "github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type PageUpdateInput struct {
	ID           int64   `db:"id"`
	Name         string  `db:"name"`
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

func (u *service) Update(ctx context.Context, in PageUpdateInput) error {
	updatedAt := time.Now()

	return u.TrManager.MakeTransaction(ctx, func(ctx context.Context) error {
		row, err := u.hub.Page().Read(ctx, page.PageReadInput{
			ID: in.ID,
		})
		if err != nil {
			return fmt.Errorf("page read", err.Error())
		}

		var metaID = row.MetaID
		if row.MetaID != nil {
			err := u.hub.Meta().Update(ctx, meta2.MetaUpdateInput{
				ID:          *row.MetaID,
				Title:       in.Title,
				Keywords:    in.Keywords,
				Description: in.Description,
			})
			if err != nil {
				return err
			}
		} else {
			metaID, err = u.hub.Meta().Create(ctx, meta2.MetaCreateInput{
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

		err = u.hub.SefUrl().Update(ctx, sefurl2.SefUrlUpdateInput{
			Url:      "/" + in.SefURL,
			Path:     "/",
			Name:     in.SefURL,
			Type:     pageType,
			ObjectID: in.ID,
			Template: in.Template,
		})
		if err != nil {
			return err
		}

		err = u.hub.Page().Update(ctx, page.PageUpdateInput{
			ID:           in.ID,
			Name:         in.Name,
			Content:      in.Content,
			Type:         in.Type,
			Sort:         in.Sort,
			ShortContent: in.ShortContent,
			ImageID:      in.ImageID,
			UpdatedAT:    updatedAt.Unix(),
			MetaID:       metaID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
