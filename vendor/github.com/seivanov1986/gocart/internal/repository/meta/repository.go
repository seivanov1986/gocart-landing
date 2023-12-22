package meta

import (
	"github.com/seivanov1986/sql_client"
)

type repository struct {
	db  sql_client.DataBase
	Trx sql_client.TransactionManager
}

func New(db sql_client.DataBase, Trx sql_client.TransactionManager) *repository {
	return &repository{
		db:  db,
		Trx: Trx,
	}
}
