package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/google/wire"

	uuid2 "github.com/google/uuid"
)

type WrapUUID struct {
}

func (wd *WrapUUID) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid2.New()
		ctx := context.WithValue(r.Context(), ContextKeyRequestId, uuid)
		startTime := time.Now()
		ctx = context.WithValue(ctx, ContextKeyRequestStartAt, startTime)
		logrus.Printf("The request with UUID: %v is startd at: %s", uuid, startTime.Format(time.RFC3339))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var NewWrapUUID = wire.Struct(new(WrapUUID), "*")
