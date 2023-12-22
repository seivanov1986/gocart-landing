package category_closure

import (
	"context"
)

type row struct {
	Parent int64 `db:"parent"`
	Child  int64 `db:"child"`
	Depth  int64 `db:"depth"`
}

func (cc *repository) Create(ctx context.Context, idObject, idParent int64) error {
	trx := cc.Trx.FindTransaction(ctx)

	rows := []row{}
	err := trx.SelectContext(ctx, &rows,
		`SELECT parent, ? as child, depth+1 as depth FROM category_closure WHERE child = ? 
		UNION ALL SELECT ?, ?, 0`,
		idObject, idParent, idObject, idObject,
	)
	if err != nil {
		return err
	}

	for _, v := range rows {
		_, err = trx.NamedExecContext(
			context.TODO(),
			`INSERT INTO category_closure (parent, child, depth) VALUES (:parent, :child, :depth)`,
			v,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
