package category

import (
	"context"
)

func (u *service) Delete(ctx context.Context, IDs []int64) error {
	var result error

	for _, idCategory := range IDs {
		err := u.TrManager.MakeTransaction(ctx, func(ctx context.Context) error {
			category, err := u.hub.Category().Read(ctx, idCategory)
			if err != nil {
				return err
			}

			err = u.hub.ImageToCategory().DeleteImagesInCategory(ctx, idCategory)
			if err != nil {
				return err
			}

			err = u.hub.SefUrl().DeleteByObjectType(ctx, idCategory, 2)
			if err != nil {
				return err
			}

			if category.MetaID != nil {
				err = u.hub.Meta().Delete(ctx, *category.MetaID)
				if err != nil {
					return err
				}
			}

			err = u.hub.Category().Delete(ctx, idCategory)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			result = err
		}
	}

	return result
}
