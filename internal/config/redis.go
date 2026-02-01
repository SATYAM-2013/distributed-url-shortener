package config

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	log.Printf("✅ Connected to Redis at %s\n", addr)
	return rdb
}
