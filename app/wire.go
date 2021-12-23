//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"go-auth-flow/handlers"
	"go-auth-flow/internal/config"
	"go-auth-flow/internal/database"
	"go-auth-flow/middlewares"

	"github.com/google/wire"
)

type App struct {
	AppConfig  *config.AppConfiguration
	DB         *database.Database
	Middleware *middlewares.Middlewares
	Handlers   *handlers.Handlers
}

func GetApp(ctx context.Context) (*App, error) {
	wire.Build(
		config.NewAppConfigurationSet,
		database.NewDatabaseSet,
		middlewares.MiddlewareSet,
		handlers.HandlerSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
