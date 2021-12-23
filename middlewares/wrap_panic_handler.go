package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/google/wire"

	"github.com/sirupsen/logrus"
)

type PanicHandler struct{}

func (ph *PanicHandler) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logrus.Error(rec)
				logrus.Printf("The request with UUID: %v stack trace: %v, ", r.Context().Value(ContextKeyRequestId), string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

var NewPanicHandler = wire.Struct(new(PanicHandler), "*")
