package api

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time-tracker/entity"
)

type Middleware struct {
	logger *slog.Logger
}

func NewMiddleware(l *slog.Logger) *Middleware {
	return &Middleware{
		logger: l,
	}
}

func (m *Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := m.logger.With("request_id", uuid.NewString())
		ctx := context.WithValue(context.Background(), entity.CtxLogger{}, l)
		r = r.WithContext(ctx)
		l.Info("incoming API request", "method", r.Method, "request", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
