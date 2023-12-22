package sefurl

import (
	"context"
)

func (u *repository) DeleteByObjectType(ctx context.Context, ObjectID, Type int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM sefurl WHERE id_object=? and type=?`,
		ObjectID, Type,
	)
	return err
}
