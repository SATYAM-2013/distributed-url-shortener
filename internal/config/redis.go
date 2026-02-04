package config

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(redisURL string) *redis.Client {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse REDIS_ADDR: %v", err)
	}

	rdb := redis.NewClient(opt)

	log.Println("Connected to Redis via URL")
	return rdb
}
