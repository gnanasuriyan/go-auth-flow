package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/google/wire"

	"github.com/sirupsen/logrus"
)

type WrapRequestLogger struct {
}

func (rl *WrapRequestLogger) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if startTime, ok := r.Context().Value(ContextKeyRequestStartAt).(time.Time); !ok {
				logrus.Printf("Getting start time from the context was failed")
			} else {
				endTime := time.Now()
				logrus.Printf("The request with UUID: %v is finished at: %s. Total time taken: %d ms", r.Context().Value(ContextKeyRequestId), endTime.Format(time.RFC3339), endTime.Sub(startTime))
			}
		}()
		startTime := time.Now()
		logrus.Printf("The request with UUID: %v is startd at: %s", r.Context().Value(ContextKeyRequestId), startTime.Format(time.RFC3339))
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ContextKeyRequestStartAt, startTime)))
	})
}

var NewWrapRequestLogger = wire.Struct(new(WrapRequestLogger), "*")
