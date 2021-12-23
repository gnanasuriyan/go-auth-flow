package handlers

import "github.com/google/wire"

type Handlers struct {
	LoginHandler *LoginHandler
}

var HandlerSet = wire.NewSet(
	NewLoginHandler,
	wire.Struct(new(Handlers), "*"),
)
