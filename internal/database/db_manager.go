package database

import (
	"context"
	"database/sql/driver"
	"fmt"
	"go-auth-flow/internal/config"
	"log"
	"sync"

	"github.com/google/wire"

	"github.com/jmoiron/sqlx"

	"github.com/elgris/sqrl"
)

type IDb interface {
	ExecContext(ctx context.Context, sqlizer sqrl.Sqlizer) error
	SelectContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error
	GetContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error
}

type DB struct {
	DB *sqlx.DB
}

var NewDB = wire.NewSet(
	InitializeDatabaseConnection,
	wire.Bind(new(IDb), new(DB)),
)

var (
	once sync.Once
	db   *DB
)

func InitializeDatabaseConnection(dbConfig config.IDatabaseConfig) *DB {
	once.Do(func() {
		connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetHost(), dbConfig.GetPort(), dbConfig.GetDatabase())
		sqlxDB, err := sqlx.Open("mysql", connectStr)
		if err != nil {
			log.Fatal(err)
		}

		sqlxDB.SetMaxIdleConns(dbConfig.GetMaxIdleConnections())
		sqlxDB.SetMaxOpenConns(dbConfig.GetMaxOpenConnections())
		db = &DB{DB: sqlxDB}
	})
	return db
}

func (db *DB) ExecContext(ctx context.Context, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDb := db.DB
	_, err = sqlxDb.ExecContext(ctx, query, args...)
	return err
}

func (db *DB) SelectContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDb := db.DB
	return sqlxDb.SelectContext(ctx, dest, query, args...)
}

func (db *DB) GetContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDb := db.DB
	return sqlxDb.GetContext(ctx, dest, query, args...)
}

type IDbTx interface {
	IDb
	driver.Tx
}

type Tx struct {
	*DB
	Tx *sqlx.Tx
}

func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}
