package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	rateLimitRequests = 10          // requests
	rateLimitWindow   = time.Minute // per minute
	apiKeyHeader      = "X-API-Key"
)

func ApiKeyRateLimiter(rdb *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 1️⃣ API key must exist
			apiKey := r.Header.Get(apiKeyHeader)
			if apiKey == "" {
				http.Error(w, "missing API key", http.StatusUnauthorized)
				return
			}

			// 2️⃣ Time window (minute-based)
			window := time.Now().Unix() / int64(rateLimitWindow.Seconds())
			key := "rate_limit:" + apiKey + ":" + strconv.FormatInt(window, 10)

			ctx := context.Background()

			// 3️⃣ Atomic increment
			count, err := rdb.Incr(ctx, key).Result()
			if err != nil {
				http.Error(w, "rate limiter error", http.StatusInternalServerError)
				return
			}

			// 4️⃣ Set TTL only once
			if count == 1 {
				rdb.Expire(ctx, key, rateLimitWindow)
			}

			// 5️⃣ Enforce limit
			if count > rateLimitRequests {
				http.Error(w, "too many requests", http.StatusTooManyRequests)
				return
			}

			// 6️⃣ Allow request
			next.ServeHTTP(w, r)
		})
	}
}
