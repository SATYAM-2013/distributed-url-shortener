package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {

	// Priority 1: REDIS_URL (Cloud providers like Upstash)
	if redisURL := os.Getenv("REDIS_URL"); redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("❌ Failed to parse REDIS_URL: %v", err)
		}

		client := redis.NewClient(opt)

		testConnection(client)
		log.Println("✅ Connected to Redis via REDIS_URL")

		return client
	}

	// Priority 2: REDIS_ADDR (Docker/local network)
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "127.0.0.1:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	testConnection(client)
	log.Printf("✅ Connected to Redis at %s\n", redisAddr)

	return client
}

func testConnection(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}
}
