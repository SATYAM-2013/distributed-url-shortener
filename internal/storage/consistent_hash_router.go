package storage

import (
	"hash/fnv"

	"github.com/redis/go-redis/v9"
)

type ConsistentHashRouter struct {
	shards []*redis.Client
}

func NewConsistentHashRouter(shards []*redis.Client) *ConsistentHashRouter {
	return &ConsistentHashRouter{shards: shards}
}

func (c *ConsistentHashRouter) ClientForKey(key string) *redis.Client {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	index := int(h.Sum32()) % len(c.shards)
	return c.shards[index]
}
