// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/google/wire"
	"oauth-study/handlers"
	"oauth-study/internal/config"
	"oauth-study/internal/database"
	"oauth-study/middlewares"
)

// Injectors from wire.go:

func GetApp(ctx context.Context) (*App, error) {
	loginHandler := handlers.LoginHandler{}
	appConfig := config.GetConfig()
	db := database.InitializeDatabaseConnection(appConfig)
	panicHandler := &middlewares.PanicHandler{}
	wrapRequestLogger := &middlewares.WrapRequestLogger{}
	wrapUUID := &middlewares.WrapUUID{}
	middleware := &middlewares.Middleware{
		PanicHandler:      panicHandler,
		WrapRequestLogger: wrapRequestLogger,
		WrapUUID:          wrapUUID,
	}
	app := &App{
		LoginHandler: loginHandler,
		AppConfig:    appConfig,
		DB:           db,
		Middleware:   middleware,
	}
	return app, nil
}

// wire.go:

type App struct {
	LoginHandler handlers.ILoginHandler
	AppConfig    *config.AppConfig
	DB           *database.DB
	Middleware   *middlewares.Middleware
}

var handlerSet = wire.NewSet(handlers.NewLoginHandler)

var middlewareSet = wire.NewSet(wire.Struct(new(middlewares.PanicHandler), "*"), wire.Struct(new(middlewares.WrapRequestLogger), "*"), wire.Struct(new(middlewares.WrapUUID), "*"))
