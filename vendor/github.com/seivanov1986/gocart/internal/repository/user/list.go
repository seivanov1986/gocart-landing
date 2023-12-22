package user

import (
	"context"
)

type UserListInput struct {
	Page int64
}

type UserListOut struct {
	List  []UserListRow
	Total int64
}

type UserListRow struct {
	ID    int64   `db:"id" json:"id"`
	Login string  `db:"login" json:"login"`
	Email *string `db:"email" json:"email"`
}

func (r *repository) List(ctx context.Context, in UserListInput) (*UserListOut, error) {
	pageRows := []UserListRow{}
	err := r.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, login, email FROM user LIMIT ?, ?`,
		in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = r.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM user`)
	if err != nil {
		return nil, err
	}

	return &UserListOut{
		List:  pageRows,
		Total: total,
	}, nil
}
