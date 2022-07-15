package cache

import (
	"context"
	"errors"
	"time"

	gatewayCache "xti-gateway-graph-go/cache"

	"github.com/patrickmn/go-cache"
)

type AppCache struct {
	client *cache.Cache
}

var (
	// DefaultCache is the default cache.
	DefaultExpiration time.Duration = 0
	CleanUpInterval   time.Duration = 10 * time.Minute

	// ErrItemExpired is returned in Cache.Get when the item found in the cache
	// has expired.
	ErrItemExpired error = errors.New("item has expired")
	// ErrKeyNotFound is returned in Cache.Get and Cache.Delete when the
	// provided key could not be found in cache.
	ErrKeyNotFound error = errors.New("key not found in cache")
)

func NewAppCache(cache *cache.Cache) *AppCache {
	return &AppCache{
		client: cache,
	}
}

func (c *AppCache) Context(ctx context.Context) gatewayCache.Cache {
	return gatewayCache.NewCache()
}

func (c *AppCache) Put(key string, data interface{}, expiration time.Duration) error {
	return c.client.Add(key, data, expiration)
}

func (c *AppCache) Get(key string) (interface{}, time.Time, error) {
	if _, ok := c.client.Get(key); !ok {
		return nil, time.Time{}, ErrKeyNotFound
	}
	cache, duration, _ := c.client.GetWithExpiration(key)
	return cache, duration, nil
}

func (c *AppCache) Delete(key string) error {
	if _, ok := c.client.Get(key); !ok {
		return ErrKeyNotFound
	}
	c.client.Delete(key)
	return nil
}

func (c *AppCache) Expired(expiration time.Time) bool {
	if expiration.IsZero() {
		return true
	}

	return time.Now().UnixNano() > expiration.UnixNano()
}
