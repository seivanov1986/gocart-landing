package sql_client

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type key string

const (
	trx key = "trx"

	migrationsDirectory = "."
)

func (d *DataBaseImpl) NewTransaction() (*sqlxTransaction, error) {
	tx, _ := d.DB.Beginx()
	return &sqlxTransaction{tx}, nil
}

func (d *DataBaseImpl) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.DB.SelectContext(ctx, dest, query, args...)
}

func (d *DataBaseImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.DB.ExecContext(ctx, query, args...)
}

func (d *DataBaseImpl) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.DB.NamedExecContext(ctx, query, arg)
}

func (d *DataBaseImpl) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.DB.GetContext(ctx, dest, query, args...)
}

func (d *DataBaseImpl) DeleteIn(ctx context.Context, query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = d.DB.ExecContext(ctx, query, inArgs...)
	return err
}

func (d *DataBaseImpl) RunMigrations(l goose.Logger, migrationFiles fs.FS) error {
	goose.SetBaseFS(migrationFiles)
	goose.SetDialect(d.DB.DriverName())
	goose.SetLogger(l)
	if err := goose.Up(d.DB.DB, migrationsDirectory); err != nil {
		return fmt.Errorf("failure to perform migrations: %v", err)
	}
	return nil
}

func (d *DataBaseImpl) Close() error {
	return d.DB.Close()
}

func (t *sqlxTransaction) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.TX.SelectContext(ctx, dest, query, args...)
}

func (t *sqlxTransaction) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return t.TX.ExecContext(ctx, query, args...)
}

func (t *sqlxTransaction) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return t.TX.NamedExecContext(ctx, query, arg)
}

func (t *sqlxTransaction) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.TX.GetContext(ctx, dest, query, args...)
}

func (t *sqlxTransaction) DeleteIn(ctx context.Context, query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = t.TX.ExecContext(ctx, query, inArgs...)
	return err
}

func (tr *transactionManager) MakeTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	transaction, err := tr.db.NewTransaction()
	if err != nil {
		return err
	}

	err = fn(context.WithValue(ctx, trx, transaction))
	if err != nil {
		transaction.TX.Rollback()
		return err
	}

	return transaction.TX.Commit()
}

func (tr *transactionManager) FindTransaction(ctx context.Context) *sqlxTransaction {
	transaction := ctx.Value(trx)
	result, ok := transaction.(*sqlxTransaction)
	if !ok {
		return nil
	}

	return result
}
