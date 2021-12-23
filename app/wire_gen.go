// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"context"
	"go-auth-flow/internal/config"
	"go-auth-flow/internal/database"
	"go-auth-flow/middlewares"
)

// Injectors from wire.go:

func GetApp(ctx context.Context) (*App, error) {
	configuration := config.GetConfiguration()
	appConfiguration := &config.AppConfiguration{
		Configuration: configuration,
	}
	dbDependencies := database.DBDependencies{
		DatabaseConfig: appConfiguration,
	}
	db := database.InitializeDatabaseConnection(dbDependencies)
	databaseDatabase := &database.Database{
		SqlxDB: db,
	}
	panicHandler := &middlewares.PanicHandler{}
	wrapRequestLogger := &middlewares.WrapRequestLogger{}
	wrapUUID := &middlewares.WrapUUID{}
	middlewaresMiddlewares := &middlewares.Middlewares{
		PanicHandler:      panicHandler,
		WrapRequestLogger: wrapRequestLogger,
		WrapUUID:          wrapUUID,
	}
	app := &App{
		AppConfig:  appConfiguration,
		DB:         databaseDatabase,
		Middleware: middlewaresMiddlewares,
	}
	return app, nil
}

// wire.go:

type App struct {
	AppConfig  *config.AppConfiguration
	DB         *database.Database
	Middleware *middlewares.Middlewares
}
