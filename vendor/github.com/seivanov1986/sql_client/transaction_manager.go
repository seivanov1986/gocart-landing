package sql_client

type transactionManager struct {
	db DataBase
}

func NewTransactionManager(db DataBase) *transactionManager {
	return &transactionManager{db: db}
}
