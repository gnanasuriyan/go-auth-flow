//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"go-auth-flow/internal/config"

	"github.com/google/wire"
)

type App struct {
	AppConfig *config.AppConfiguration
	//DB        *database.DB
	//Middleware   *middlewares.Middleware
	//LoginHandler handlers.ILoginHandler
}

//var handlerSet = wire.NewSet(
//	handlers.NewLoginHandler,
//)

//var middlewareSet = wire.NewSet(
//	wire.Struct(new(middlewares.PanicHandler), "*"),
//	wire.Struct(new(middlewares.WrapRequestLogger), "*"),
//	wire.Struct(new(middlewares.WrapUUID), "*"),
//)

func GetApp(ctx context.Context) (*App, error) {
	wire.Build(
		config.NewAppConfiguration,
		//config.NewDB,
		//config.GetConfiguration,
		//wire.Bind(new(config.IDatabaseConfig), new(config.Configuration)),
		//handlerSet,
		//middlewareSet,
		//wire.Struct(new(middlewares.Middleware), "*"),
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
