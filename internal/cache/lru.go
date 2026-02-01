package cache

import (
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Item struct {
	URL       string
	ExpiresAt time.Time
}

type LRUCache struct {
	cache *lru.Cache[string, Item]
}

func NewLRUCache(size int) (*LRUCache, error) {
	c, err := lru.New[string, Item](size)
	if err != nil {
		return nil, err
	}
	return &LRUCache{cache: c}, nil
}

func (c *LRUCache) Get(key string) (string, bool) {
	item, ok := c.cache.Get(key)
	if !ok || time.Now().After(item.ExpiresAt) {
		c.cache.Remove(key)
		return "", false
	}
	return item.URL, true
}

func (c *LRUCache) Set(key, url string, ttl time.Duration) {
	c.cache.Add(key, Item{
		URL:       url,
		ExpiresAt: time.Now().Add(ttl),
	})
}
