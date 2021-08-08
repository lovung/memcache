package memcache

import (
	"sync"
	"time"
)

const (
	defaultTimeout           = 30 * time.Minute
	defaultNumberOfShard int = 1 << 8
	mask                     = 255 // 0xFF
)

// Cache represent the cache interface
type Cache interface {
	Set(key string, value interface{})
	SetUntil(key string, value interface{}, d time.Duration)
	Get(key string) interface{}
	Take(key string) interface{}
	Del(key string)
}

type cache struct {
	shards [defaultNumberOfShard]shard
	mu     [defaultNumberOfShard]sync.Mutex
}

type item struct {
	expiredAt time.Time
	value     interface{}
}

// NewCache creates a memory cache
func NewCache() Cache {
	newCache := cache{}
	for i := range newCache.shards {
		newCache.shards[i].hashmap = make(map[string]item)
	}
	return &newCache
}

func (c *cache) Set(key string, value interface{}) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	defer c.mu[shardIdx].Unlock()
	c.shards[shardIdx].set(key, value, defaultTimeout)
}

func (c *cache) SetUntil(key string, value interface{}, d time.Duration) {
	if d <= 0 {
		return
	}
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	c.shards[shardIdx].set(key, value, d)
	c.mu[shardIdx].Unlock()
}

func (c *cache) Get(key string) (value interface{}) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	defer c.mu[shardIdx].Unlock()
	return c.shards[shardIdx].get(key)
}

func (c *cache) Take(key string) (value interface{}) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	defer c.mu[shardIdx].Unlock()
	return c.shards[shardIdx].take(key)
}

func (c *cache) Del(key string) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	c.shards[shardIdx].delete(key)
	c.mu[shardIdx].Unlock()
}
