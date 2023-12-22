package sefurl

import (
	"context"
)

type SefUrlListInput struct {
	Page int64
}

type SefUrlListOut struct {
	List  []SefUrlListRow
	Total int64
}

type SefUrlListRow struct {
	ID       int64  `db:"id" json:"id"`
	Url      string `db:"url" json:"url"`
	Path     string `db:"path" json:"path"`
	Name     string `db:"name" json:"name"`
	Type     int64  `db:"type" json:"type"`
	IdObject int64  `db:"id_object" json:"id_object"`
}

func (i *repository) List(ctx context.Context, in SefUrlListInput) (*SefUrlListOut, error) {
	pageRows := []SefUrlListRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, url, path, name, type, id_object FROM sefurl LIMIT ?, ?`,
		in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM sefurl`)
	if err != nil {
		return nil, err
	}

	return &SefUrlListOut{
		List:  pageRows,
		Total: total,
	}, nil
}
