package goCache

import (
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	cacheConfig "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/metrics"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/value-wrapper"
	"sync/atomic"
	"unsafe"
)

// GoCache go-cache缓存
type GoCache struct {
	cacheName    string
	cache        *cache.Cache
	cacheMetrics *metrics.CacheMetrics
}

// NewGoCache 创建一个新的GoCache实例
func NewGoCache(cacheConfig *cacheConfig.GoCacheConfig) *GoCache {
	cacheMaxCount := int64(0)
	if viper.GetInt64("config.cache.cache_max_count") == 0 {
		cacheMaxCount = 1000
	}
	return &GoCache{
		cacheName:    cacheConfig.CacheName,
		cache:        cache.New(cacheConfig.ExpireTime, cacheConfig.CleanTime),
		cacheMetrics: metrics.NewCacheMetrics(cacheMaxCount),
	}
}

func (c *GoCache) Set(key string, value interface{}) {
	//if c.cache.ItemCount() >= int(viper.GetInt64("config.cache.cache_max_count")) {
	//	log.Printf("cache is full, key: %s", key)
	//	return
	//}
	c.cache.Set(key, valueWrapper.NewValueWrapper(value), cache.DefaultExpiration)
}

func (c *GoCache) Get(key string) *valueWrapper.ValueWrapper {
	defer func() {
		atomic.AddInt64(&c.cacheMetrics.CacheQueryCount, 1)
	}()
	item, found := c.cache.Get(key)
	if found {
		vw := item.(*valueWrapper.ValueWrapper)
		atomic.AddInt64(&c.cacheMetrics.CacheHitCount, 1)
		vw.UpdateCacheValueMetrics()
		vw.UpdateAccessTime()
		vw.UpdateWriteTime()
		return vw
	} else {
		atomic.AddInt64(&c.cacheMetrics.CacheMissCount, 1)
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
func (c *GoCache) GetCacheValuesStatus() []*metrics.CacheValueMetrics {
	return nil
}

// GetCacheMetrics 统计缓存状态
func (c *GoCache) GetCacheMetrics() *metrics.CacheMetrics {
	metrics := c.cacheMetrics
	metrics.CacheCurrentKeyCount = int64(c.cache.ItemCount())
	metrics.CacheSize = int64(unsafe.Sizeof(c.cache))
	if metrics.CacheQueryCount == 0 {
		metrics.CacheHitRate = 0
		metrics.CacheMissRate = 0
		return metrics
	}
	metrics.CacheHitRate = (float32(metrics.CacheHitCount) / float32(metrics.CacheQueryCount)) * float32(100)
	metrics.CacheMissRate = float32(metrics.CacheMissCount) / float32(metrics.CacheQueryCount) * float32(100)
	return metrics
}
