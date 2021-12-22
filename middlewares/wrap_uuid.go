package middlewares

import (
	"context"
	"net/http"

	uuid2 "github.com/google/uuid"
)

type WrapUUID struct {
}

func (wd *WrapUUID) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid2.New()
		ctx := context.WithValue(r.Context(), ContextKeyRequestId, uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
