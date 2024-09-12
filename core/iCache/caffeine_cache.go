package iCache

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type CaffeineCache struct {
	name  string
	cache *cache.Cache
	lock  sync.RWMutex
}

// NewCaffeineCache 创建一个新的CaffeineCache实例
func NewCaffeineCache(name string, config *CacheConfig) *CaffeineCache {
	return &CaffeineCache{
		name:  name,
		cache: cache.New(config.ExpireAfterWrite, config.ExpireAfterAccess),
	}
}

func (c *CaffeineCache) Get(key string) *ValueWrapper {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if item, found := c.cache.Get(key); found {
		vw := item.(*ValueWrapper)
		vw.UpdateAccessTime()
		return vw
	}
	return nil
}

func (c *CaffeineCache) Put(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Set(key, NewValueWrapper(value), cache.DefaultExpiration)
}

func (c *CaffeineCache) GetValues() []*ValueWrapper {
	c.lock.RLock()
	defer c.lock.RUnlock()
	var values []*ValueWrapper
	for _, item := range c.cache.Items() {
		values = append(values, item.Object.(*ValueWrapper))
	}
	return values
}

func (c *CaffeineCache) GetKeys() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	var keys []string
	for key := range c.cache.Items() {
		keys = append(keys, key)
	}
	return keys
}

func (c *CaffeineCache) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cache.ItemCount()
}

func (c *CaffeineCache) GetName() string {
	return c.name
}

func (c *CaffeineCache) CalculateMemoryUsage() float64 {
	// This is a simplified version and does not calculate actual memory usage
	return float64(c.Size())
}

func (c *CaffeineCache) GetCacheStatus() CacheStats {
	// This is a simplified version and does not provide real cache statistics
	return CacheStats{}
}

// CacheConfig 缓存配置
type CacheConfig struct {
	ExpireAfterWrite  time.Duration
	ExpireAfterAccess time.Duration
}
