package database

import (
	"context"
	"fmt"
	"go-auth-flow/internal/config"
	"log"
	"sync"

	"github.com/google/wire"

	"github.com/jmoiron/sqlx"

	"github.com/elgris/sqrl"
)

type IDatabase interface {
	ExecContext(ctx context.Context, sqlizer sqrl.Sqlizer) error
	SelectContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error
	GetContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error
}

type DBDependencies struct {
	DatabaseConfig config.IDatabaseConfig
}

type Database struct {
	SqlxDB *sqlx.DB
}

var NewDatabase = wire.NewSet(
	wire.Struct(new(DBDependencies), "*"),
	InitializeDatabaseConnection,
	wire.Struct(new(Database), "*"),
	wire.Bind(new(IDatabase), new(*Database)),
)

var (
	once   sync.Once
	sqlxDB *sqlx.DB
)

func InitializeDatabaseConnection(dep DBDependencies) *sqlx.DB {
	once.Do(func() {
		dbConfig := dep.DatabaseConfig
		connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetHost(), dbConfig.GetPort(), dbConfig.GetDatabase())
		sqlxDB, err := sqlx.Open("mysql", connectStr)
		if err != nil {
			log.Fatal(err)
		}
		sqlxDB.SetMaxIdleConns(dbConfig.GetMaxIdleConnections())
		sqlxDB.SetMaxOpenConns(dbConfig.GetMaxOpenConnections())
	})
	return sqlxDB
}

func (db *Database) ExecContext(ctx context.Context, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDb := db.SqlxDB
	_, err = sqlxDb.ExecContext(ctx, query, args...)
	return err
}

func (db *Database) SelectContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDb := db.SqlxDB
	return sqlxDb.SelectContext(ctx, dest, query, args...)
}

func (db *Database) GetContext(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDb := db.SqlxDB
	return sqlxDb.GetContext(ctx, dest, query, args...)
}
