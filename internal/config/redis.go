package config

import (
	"log"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_URL")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	// If full URL (Upstash)
	if strings.HasPrefix(addr, "redis://") || strings.HasPrefix(addr, "rediss://") {
		opt, err := redis.ParseURL(addr)
		if err != nil {
			log.Fatalf("Failed to parse REDIS_URL: %v", err)
		}
		return redis.NewClient(opt)
	}

	// Local redis
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
