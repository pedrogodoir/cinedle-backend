package cache

import (
	"sync"
	"time"
)

// Requires go >= 1.18
// This implementation uses Go Generics

// CacheItem represents an item stored in the cache with its associated TTL.
type CacheItem[T any] struct {
	value  T
	expiry time.Time
}

// Cache represents an in-memory key-value store with expiry support.
type Cache[K comparable, T any] struct {
	data map[K]CacheItem[T] // stores cache items
	ttl  time.Duration      // default time-to-live for items
	mu   sync.RWMutex       // managing concurrent access
}

// NewCache creates and initializes a new Cache instance.
func NewCache[K comparable, T any](ttl time.Duration) *Cache[K, T] {
	return &Cache[K, T]{
		data: make(map[K]CacheItem[T], 10),
		ttl:  ttl,
	}
}

// Set adds or updates a key-value pair in the cache with the given TTL.
func (c *Cache[K, T]) Set(key K, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = CacheItem[T]{
		value:  value,
		expiry: time.Now().Add(c.ttl),
	}
}

func zeroVal[T any]() T {
	var zero T
	return zero
}

// Get retrieves the value associated with the given key from the cache.
// It also checks for expiry and removes expired items.
func (c *Cache[K, T]) Get(key K) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.data[key]
	if !ok {
		return zeroVal[T](), false
	}
	// item found - check for expiry
	if item.expiry.Before(time.Now()) {
		// remove entry from cache if time is beyond the expiry
		delete(c.data, key)
		return zeroVal[T](), false
	}
	// refresh expiry
	item.expiry = time.Now().Add(c.ttl)

	return item.value, true
}

// Delete removes a key-value pair from the cache.
func (c *Cache[K, T]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Clear removes all key-value pairs from the cache.
func (c *Cache[K, T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[K]CacheItem[T])
}
