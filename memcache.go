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

type Stringer interface {
	~string
}

// Cache represent the memory cache
type Cache[K Stringer, V any] struct {
	shards [defaultNumberOfShard]shard[K, V]
	mu     [defaultNumberOfShard]sync.Mutex
}

type item[V any] struct {
	expiredAt time.Time
	value     V
}

// NewCache creates a memory cache
func NewCache[K Stringer, V any]() Cache[K,V] {
	newCache := Cache[K, V]{}
	for i := range newCache.shards {
		newCache.shards[i].hashmap = make(map[K]item[V])
	}
	return &newCache
}

// Set the value with the key
func (c *Cache[K Stringer, V any]) Set(key K, value V) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	defer c.mu[shardIdx].Unlock()
	c.shards[shardIdx].set(key, value, defaultTimeout)
}

func (c *Cache[K Stringer, V any]) SetUntil(key K, value V, d time.Duration) {
	if d <= 0 {
		return
	}
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	c.shards[shardIdx].set(key, value, d)
	c.mu[shardIdx].Unlock()
}

func (c *Cache[K Stringer, V any]) Get(key K) (value V) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	defer c.mu[shardIdx].Unlock()
	return c.shards[shardIdx].get(key)
}

func (c *Cache[K Stringer, V any]) Take(key K) (value V) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	defer c.mu[shardIdx].Unlock()
	return c.shards[shardIdx].take(key)
}

func (c *Cache[K Stringer, V any]) Del(key K) {
	hashKey := hashString(key)
	shardIdx := hashKey & mask
	c.mu[shardIdx].Lock()
	c.shards[shardIdx].delete(key)
	c.mu[shardIdx].Unlock()
}
