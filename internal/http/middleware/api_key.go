package middleware

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type APIKeyMiddleware struct {
	rdb *redis.Client
}

func NewAPIKeyMiddleware(rdb *redis.Client) *APIKeyMiddleware {
	return &APIKeyMiddleware{rdb: rdb}
}

func (a *APIKeyMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "missing API key", http.StatusUnauthorized)
			return
		}

		ctx := context.Background()
		exists, err := a.rdb.Exists(ctx, "api_key:"+apiKey).Result()
		if err != nil || exists == 0 {
			http.Error(w, "invalid API key", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
