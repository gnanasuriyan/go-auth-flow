package middlewares

import "github.com/google/wire"

type Middleware struct {
	PanicHandler      *PanicHandler
	WrapRequestLogger *WrapRequestLogger
	WrapUUID          *WrapUUID
}

var MiddlewareSet = wire.NewSet(
	NewPanicHandler,
	NewWrapRequestLogger,
	NewWrapUUID,
	wire.Struct(new(Middleware), "*"),
)
