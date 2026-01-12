package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type ctxKey string

const RequestIdKey ctxKey = "request_id"

func (h *MealPlanHandler) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqId := uuid.New().String()
		ctx := context.WithValue(r.Context(), RequestIdKey, reqId)
		r = r.WithContext(ctx)
		w.Header().Set("X-Request-Id", reqId)
		next.ServeHTTP(w, r)
		slog.Info("HTTP Request",
			"id", reqId,
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}
