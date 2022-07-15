package cache

import (
	"context"
	"errors"
	"time"

	gatewayCache "xti-gateway-graph-go/cache"

	"github.com/gomodule/redigo/redis"
)

type AppCache struct {
	client *redis.Pool
}

var (
	// DefaultCache is the default cache.
	DefaultExpiration time.Duration = 0
	CleanUpInterval   time.Duration = 10 * time.Minute
	DefaultPool       int           = 20
	DefaultHost       string        = "127.0.0.1"
	DefaultPort       int           = 6379

	// ErrItemExpired is returned in Cache.Get when the item found in the cache
	// has expired.
	ErrItemExpired error = errors.New("item has expired")
	// ErrKeyNotFound is returned in Cache.Get and Cache.Delete when the
	// provided key could not be found in cache.
	ErrKeyNotFound error = errors.New("key not found in cache")
)

func NewAppCache(cache *redis.Pool) *AppCache {
	return &AppCache{
		client: cache,
	}
}

func (c *AppCache) Context(ctx context.Context) gatewayCache.Cache {
	return gatewayCache.NewCache()
}

func (c *AppCache) Put(key string, data interface{}, expiration time.Duration) error {
	conn := c.client.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, data, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (c *AppCache) Get(key string) (interface{}, time.Time, error) {
	conn := c.client.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, time.Time{}, err
	}
	return data, time.Time{}, nil
}

func (c *AppCache) Delete(key string) error {
	if _, _, ok := c.Get(key); ok != nil {
		return ErrKeyNotFound
	}
	conn := c.client.Get()
	defer conn.Close()
	_, err := redis.Bytes(conn.Do("DELETE", key))
	if err != nil {
		return err
	}
	return nil
}

func (c *AppCache) Expired(expiration time.Time) bool {
	if expiration.IsZero() {
		return true
	}

	return time.Now().UnixNano() > expiration.UnixNano()
}
