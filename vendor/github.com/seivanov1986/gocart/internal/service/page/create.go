package page

import (
	"context"
	"fmt"
	"time"

	"github.com/seivanov1986/gocart/internal/repository/meta"
	"github.com/seivanov1986/gocart/internal/repository/page"
	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type PageCreateIn struct {
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

func (u *service) Create(ctx context.Context, in PageCreateIn) (*int64, error) {
	var pageID *int64
	createdAt := time.Now()
	updatedAt := createdAt

	err := u.TrManager.MakeTransaction(ctx, func(ctx context.Context) error {
		idMeta, err := u.hub.Meta().Create(ctx, meta.MetaCreateInput{
			Title:       in.Title,
			Keywords:    in.Keywords,
			Description: in.Description,
		})
		if err != nil {
			return err
		}

		pageID, err = u.hub.Page().Create(ctx, page.PageCreateInput{
			Name:         in.Name,
			Content:      in.Content,
			MetaID:       idMeta,
			Type:         in.Type,
			Sort:         in.Sort,
			ShortContent: in.ShortContent,
			ImageID:      in.ImageID,
			CreatedAT:    createdAt.Unix(),
			UpdatedAT:    updatedAt.Unix(),
		})
		if err != nil {
			return err
		}

		_, err = u.hub.SefUrl().Create(ctx, sefurl.SefUrlCreateInput{
			Url:       "/" + in.SefURL,
			Path:      "/",
			Name:      in.SefURL,
			Type:      pageType,
			ObjectID:  *pageID,
			Template:  in.Template,
			CreatedAt: createdAt.Unix(),
			UpdatedAt: updatedAt.Unix(),
		})
		if err != nil {
			return fmt.Errorf("SefUrl", err.Error())
		}

		return nil
	})

	return pageID, err
}
