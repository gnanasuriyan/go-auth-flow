package middlewares

import "github.com/google/wire"

type Middlewares struct {
	PanicHandler      *PanicHandler
	WrapRequestLogger *WrapRequestLogger
	WrapUUID          *WrapUUID
}

var MiddlewareSet = wire.NewSet(
	NewPanicHandler,
	NewWrapRequestLogger,
	NewWrapUUID,
	wire.Struct(new(Middlewares), "*"),
)
