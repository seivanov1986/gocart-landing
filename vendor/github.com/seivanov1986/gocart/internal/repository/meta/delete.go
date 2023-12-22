package meta

import (
	"context"
)

func (u *repository) Delete(ctx context.Context, idMeta int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM meta WHERE id=?`,
		idMeta,
	)
	return err
}
