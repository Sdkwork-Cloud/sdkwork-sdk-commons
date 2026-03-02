package common

import "sync"

type CacheStore struct {
	config *CacheConfig
	cache  map[string]struct {
		value    interface{}
		timestamp int64
	}
	mu   sync.Mutex
}

func NewCacheStore(config *CacheConfig) *CacheStore {
	if config == nil {
		config = &DefaultCacheConfig
	}
	return &CacheStore{
		config: config,
		cache:  make(map[string]struct {
			value    interface{}
			timestamp int64
		}),
	}
}

func (c *CacheStore) Get(key string) interface{} {
	if !c.config.Enabled {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, ok := c.cache[key]; ok {
		if time.Now().UnixMilli()-item.timestamp < c.config.TTL {
			return item.value
		}
		delete(c.cache, key)
	}
	return nil
}

func (c *CacheStore) Set(key string, value interface{}) {
	if !c.config.Enabled {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.cache) >= c.config.MaxSize {
		c.evictOldest()
	}
	c.cache[key] = struct {
		value    interface{}
		timestamp int64
	}{value, time.Now().UnixMilli()}
}

func (c *CacheStore) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}

func (c *CacheStore) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache = make(map[string]struct {
		value    interface{}
		timestamp int64
	}{})
}

func (c *CacheStore) evictOldest() {
	var oldestKey string
	var oldestTime int64 = math.MaxInt64
	for key, item := range c.cache {
		if item.timestamp < oldestTime {
			oldestTime = item.timestamp
			oldestKey = key
		}
	}
	if oldestKey != "" {
		delete(c.cache, oldestKey)
	}
}

import "math"
import "time"
