package storage

import "github.com/redis/go-redis/v9"

type ShardRouter interface {
	ClientForKey(key string) *redis.Client
}
