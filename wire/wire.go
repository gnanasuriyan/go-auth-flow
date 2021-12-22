//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"oauth-study/handlers"
	"oauth-study/internal/config"
	"oauth-study/internal/database"
	"oauth-study/middlewares"

	"github.com/google/wire"
)

type App struct {
	LoginHandler handlers.ILoginHandler
	AppConfig    *config.AppConfig
	DB           *database.DB
	Middleware   *middlewares.Middleware
}

var handlerSet = wire.NewSet(
	handlers.NewLoginHandler,
)

var middlewareSet = wire.NewSet(
	wire.Struct(new(middlewares.PanicHandler), "*"),
	wire.Struct(new(middlewares.WrapRequestLogger), "*"),
	wire.Struct(new(middlewares.WrapUUID), "*"),
)

func GetApp(ctx context.Context) (*App, error) {
	wire.Build(
		config.GetConfig,
		database.InitializeDatabaseConnection,
		handlerSet,
		middlewareSet,
		wire.Struct(new(middlewares.Middleware), "*"),
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
