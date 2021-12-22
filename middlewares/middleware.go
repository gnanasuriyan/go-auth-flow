package middlewares

type Middleware struct {
	PanicHandler      *PanicHandler
	WrapRequestLogger *WrapRequestLogger
	WrapUUID          *WrapUUID
}
