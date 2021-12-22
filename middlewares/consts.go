package middlewares

type ContextKey string

const (
	ContextKeyRequestId      ContextKey = "request_id"
	ContextKeyRequestStartAt ContextKey = "request_start_at"
)
