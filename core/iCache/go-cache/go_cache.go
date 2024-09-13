package go_cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/value-wrapper"
)

// GoCache go-cache缓存
type GoCache struct {
	name  string
	cache *cache.Cache
}

// NewGoCache 创建一个新的GoCache实例
func NewGoCache(cacheConfig *config.GoCacheConfig) *GoCache {
	return &GoCache{
		name:  cacheConfig.CacheName,
		cache: cache.New(cacheConfig.ExpireTime, cacheConfig.CleanTime),
	}
}

func (c *GoCache) Set(key string, value interface{}) {
	c.cache.Set(key, value_wrapper.NewValueWrapper(value), cache.DefaultExpiration)
}

func (c *GoCache) Get(key string) *value_wrapper.ValueWrapper {
	if item, found := c.cache.Get(key); found {
		vw := item.(*value_wrapper.ValueWrapper)
		vw.UpdateAccessTime()
		vw.UpdateWriteTime()
		return vw
	}
	return nil
}

func (c *GoCache) GetValues() []*value_wrapper.ValueWrapper {
	var values []*value_wrapper.ValueWrapper
	for _, item := range c.cache.Items() {
		values = append(values, item.Object.(*value_wrapper.ValueWrapper))
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

func (c *GoCache) GetCacheStatus() iCache.CacheStats {
	// This is a simplified version and does not provide real cache statistics
	return iCache.CacheStats{}
}