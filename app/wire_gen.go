// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"go-auth-flow/handlers"
	"go-auth-flow/internal/config"
	"go-auth-flow/internal/database"
	"go-auth-flow/middlewares"
)

// Injectors from wire.go:

func GetApp() *App {
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
	loginHandler := &handlers.LoginHandler{}
	handlersHandlers := &handlers.Handlers{
		LoginHandler: loginHandler,
	}
	app := &App{
		AppConfig:   appConfiguration,
		DB:          databaseDatabase,
		Middlewares: middlewaresMiddlewares,
		Handlers:    handlersHandlers,
	}
	return app
}

// wire.go:

type App struct {
	AppConfig   *config.AppConfiguration
	DB          *database.Database
	Middlewares *middlewares.Middlewares
	Handlers    *handlers.Handlers
}
