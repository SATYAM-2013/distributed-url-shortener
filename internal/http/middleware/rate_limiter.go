package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	rdb    *redis.Client
	limit  int
	window time.Duration
}

func NewRateLimiter(rdb *redis.Client, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		rdb:    rdb,
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	ctx := context.Background()
	redisKey := "rate_limit:" + key

	count, err := rl.rdb.Incr(ctx, redisKey).Result()
	if err != nil {
		return false
	}

	if count == 1 {
		rl.rdb.Expire(ctx, redisKey, rl.window)
	}

	return count <= int64(rl.limit)
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		if !rl.Allow(apiKey) {
			http.Error(w, "rate limit error", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
