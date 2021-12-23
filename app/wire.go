//go:build wireinject
// +build wireinject

package app

import (
	"go-auth-flow/handlers"
	"go-auth-flow/middlewares"

	"github.com/google/wire"
)

type App struct {
	//AppConfig  *config.AppConfiguration
	//DB         *database.Database
	Middlewares *middlewares.Middlewares
	Handlers    *handlers.Handlers
}

func GetApp() *App {
	wire.Build(
		//config.NewAppConfigurationSet,
		//database.NewDatabaseSet,
		middlewares.MiddlewareSet,
		handlers.HandlerSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}
}
