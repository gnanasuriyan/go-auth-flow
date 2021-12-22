package database

import (
	"context"
	"database/sql/driver"
	"fmt"
	"log"
	"oauth-study/internal/config"
	"sync"

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

var once sync.Once
var db *DB

func InitializeDatabaseConnection(appConfig *config.AppConfig) *DB {
	once.Do(func() {
		connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", appConfig.DB.Username, appConfig.DB.Password, appConfig.DB.Host, appConfig.DB.Port, appConfig.DB.Database)
		sqlxDB, err := sqlx.Open("mysql", connectStr)
		if err != nil {
			log.Fatal(err)
		}
		sqlxDB.SetMaxIdleConns(appConfig.DB.MaxIdleConnections)
		sqlxDB.SetMaxOpenConns(appConfig.DB.MaxOpenConnections)
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
