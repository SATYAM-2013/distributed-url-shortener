package service

import (
	"context"
	"time"

	"distributed-url-shortener/internal/cache"
	"distributed-url-shortener/internal/metrics"

	"github.com/redis/go-redis/v9"
)

type ShortenerService struct {
	rdb   *redis.Client
	cache *cache.LRUCache
}

func NewShortenerService(rdb *redis.Client, cache *cache.LRUCache) *ShortenerService {
	return &ShortenerService{rdb: rdb, cache: cache}
}

func (s *ShortenerService) Shorten(url string) (string, error) {
	code, err := generateCode(7)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	ttl := time.Hour

	if err := s.rdb.Set(ctx, code, url, ttl).Err(); err != nil {
		metrics.RedisErrors.Inc()
		return "", ErrServiceUnavailable
	}

	s.cache.Set(code, url, ttl)
	return code, nil
}

func (s *ShortenerService) Resolve(code string) (string, error) {
	if url, ok := s.cache.Get(code); ok {
		return url, nil
	}

	ctx := context.Background()
	url, err := s.rdb.Get(ctx, code).Result()
	if err == redis.Nil {
		return "", ErrNotFound
	}
	if err != nil {
		metrics.RedisErrors.Inc()
		return "", ErrServiceUnavailable
	}

	s.cache.Set(code, url, time.Hour)
	return url, nil
}
