//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"github.com/google/wire"
	"oauth-study/handlers"
)

type App struct {
	LoginHandler handlers.LoginHandler
}

var handlerSet = wire.NewSet(
	wire.Struct(new(handlers.LoginHandler), "*"),
)

func GetApp(ctx context.Context) (*App, error) {
	wire.Build(
		handlerSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
