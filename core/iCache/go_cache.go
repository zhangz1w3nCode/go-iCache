package iCache

import (
	"github.com/patrickmn/go-cache"
)

// GoCache go-cache缓存
type GoCache struct {
	name  string
	cache *cache.Cache
}

// NewGoCache 创建一个新的GoCache实例
func NewGoCache(cacheConfig *CacheConfig) *GoCache {
	return &GoCache{
		name:  cacheConfig.CacheName,
		cache: cache.New(cacheConfig.ExpireTime, cacheConfig.CleanTime),
	}
}

func (c *GoCache) Set(key string, value interface{}) {
	c.cache.Set(key, NewValueWrapper(value), cache.DefaultExpiration)
}

func (c *GoCache) Get(key string) *ValueWrapper {
	if item, found := c.cache.Get(key); found {
		vw := item.(*ValueWrapper)
		vw.UpdateAccessTime()
		vw.UpdateWriteTime()
		return vw
	}
	return nil
}

func (c *GoCache) GetValues() []*ValueWrapper {
	var values []*ValueWrapper
	for _, item := range c.cache.Items() {
		values = append(values, item.Object.(*ValueWrapper))
	}
	return values
}

func (c *GoCache) GetKeys() []string {
	var keys []string
	for key := range c.cache.Items() {
		keys = append(keys, key)
	}
	return keys
}

func (c *GoCache) Size() int {
	return c.cache.ItemCount()
}

func (c *GoCache) GetName() string {
	return c.name
}

func (c *GoCache) CalculateMemoryUsage() float64 {
	// This is a simplified version and does not calculate actual memory usage
	return float64(c.Size())
}

func (c *GoCache) GetCacheStatus() CacheStats {
	// This is a simplified version and does not provide real cache statistics
	return CacheStats{}
}
