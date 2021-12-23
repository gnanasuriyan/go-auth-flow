package middlewares

import (
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
				logrus.Printf("The request with UUID: %v is finished at: %s. Total time taken: %v", r.Context().Value(ContextKeyRequestId), endTime.Format(time.RFC3339), endTime.Sub(startTime))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

var NewWrapRequestLogger = wire.Struct(new(WrapRequestLogger), "*")
