package goCache

import (
	"github.com/patrickmn/go-cache"
	cacheConfig "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
	cacheMetrics "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/cache-metrics"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/value-wrapper"
	"log"
	"unsafe"
)

// GoCache go-cache缓存
type GoCache struct {
	cacheName    string
	cache        *cache.Cache
	cacheMetrics *cacheMetrics.CacheMetrics
}

// NewGoCache 创建一个新的GoCache实例
func NewGoCache(cacheConfig *cacheConfig.GoCacheConfig) *GoCache {
	return &GoCache{
		cacheName:    cacheConfig.CacheName,
		cache:        cache.New(cacheConfig.ExpireTime, cacheConfig.CleanTime),
		cacheMetrics: cacheMetrics.NewCacheMetrics(cacheConfig.CacheMaxCount),
	}
}

func (c *GoCache) Set(key string, value interface{}) {
	if c.cache.ItemCount() >= int(c.cacheMetrics.CacheMaxCount) {
		log.Printf("cache is full, key: %s", key)
		return
	}
	c.cache.Set(key, valueWrapper.NewValueWrapper(value), cache.DefaultExpiration)
}

func (c *GoCache) Get(key string) *valueWrapper.ValueWrapper {
	defer func() {
		c.cacheMetrics.CacheQueryCount++
	}()
	item, found := c.cache.Get(key)
	if found {
		vw := item.(*valueWrapper.ValueWrapper)
		c.cacheMetrics.CacheHitCount++
		vw.UpdateCacheStatus()
		vw.UpdateAccessTime()
		vw.UpdateWriteTime()
		return vw
	} else {
		c.cacheMetrics.CacheMissCount++
	}
	return nil
}

func (c *GoCache) GetValues() []*valueWrapper.ValueWrapper {
	var values []*valueWrapper.ValueWrapper
	for _, item := range c.cache.Items() {
		values = append(values, item.Object.(*valueWrapper.ValueWrapper))
	}
	return values
}

func (c *GoCache) GetKeys() []string {
	keys := []string{}
	for key := range c.cache.Items() {
		keys = append(keys, key)
	}
	return keys
}

func (c *GoCache) CacheNum() int {
	return c.cache.ItemCount()
}

func (c *GoCache) GetName() string {
	return c.cacheName
}

// GetCacheValuesStatus 统计缓存值状态
func (c *GoCache) GetCacheValuesStatus() []*valueWrapper.CacheValueStatus {
	return nil
}

// GetCacheMetrics 统计缓存状态
func (c *GoCache) GetCacheMetrics() *cacheMetrics.CacheMetrics {
	metrics := c.cacheMetrics
	metrics.CacheCurrentKeyCount = int64(c.cache.ItemCount())
	metrics.CacheSize = int64(unsafe.Sizeof(c.cache))
	if metrics.CacheQueryCount == 0 {
		metrics.CacheHitRate = 0
		metrics.CacheMissRate = 0
		return metrics
	}
	metrics.CacheHitRate = float32(metrics.CacheHitCount / metrics.CacheQueryCount)
	metrics.CacheMissRate = float32(metrics.CacheMissCount / metrics.CacheQueryCount)
	return metrics
}
