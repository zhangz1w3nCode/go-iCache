package ristrettoCache

/*
import (
	cache "github.com/dgraph-io/ristretto"
	cache2 "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache"
	cacheConfig "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/value-wrapper"
)

// RistrettoCache RistrettoCache缓存
type RistrettoCache struct {
	name  string
	cache *cache.Cache
}

// NewRistrettoCache 创建一个新的RistrettoCache实例
func NewRistrettoCache(cacheConfig *cacheConfig.RistrettoCacheConfig) *RistrettoCache {

	cache, err := cache.NewCache(&cache.Config{
		NumCounters: cacheConfig.NumCounters, // 100万计数器
		MaxCost:     cacheConfig.MaxCost,     // 100MB内存预算
		BufferItems: cacheConfig.BufferItems, // 10万缓冲区项目
		Metrics:     cacheConfig.Metrics,     // 启用指标,
	})

	if err != nil {
		panic(err)
	}

	return &RistrettoCache{
		name:  cacheConfig.CacheName,
		cache: cache,
	}
}

func (c *RistrettoCache) Set(key string, value interface{}) {
	c.cache.Set(key, valueWrapper.NewValueWrapper(value), 1)
}

func (c *RistrettoCache) Get(key string) *valueWrapper.ValueWrapper {
	if item, found := c.cache.Get(key); found {
		vw := item.(*valueWrapper.ValueWrapper)
		vw.UpdateAccessTime()
		vw.UpdateWriteTime()
		return vw
	}
	return nil
}

func (c *RistrettoCache) GetValues() []*valueWrapper.ValueWrapper {
	var values []*valueWrapper.ValueWrapper
	return values
}

func (c *RistrettoCache) GetKeys() []string {
	return []string{}
}

func (c *RistrettoCache) Size() int {
	return int(c.cache.Metrics.CostAdded() - c.cache.Metrics.CostEvicted())
}

func (c *RistrettoCache) GetName() string {
	return c.name
}

func (c *RistrettoCache) CalculateMemoryUsage() float64 {
	// This is a simplified version and does not calculate actual memory usage
	return float64(c.Size())
}

func (c *RistrettoCache) GetCacheStatus() cache2.CacheStats {
	costAdd := c.cache.Metrics.CostAdded()
	costEvicted := c.cache.Metrics.CostEvicted()
	kept := c.cache.Metrics.GetsKept()
	keyAdd := c.cache.Metrics.KeysAdded()
	keysEvicted := c.cache.Metrics.KeysEvicted()
	hits := c.cache.Metrics.Hits()
	misses := c.cache.Metrics.Misses()
	updated := c.cache.Metrics.KeysUpdated()
	rejected := c.cache.Metrics.SetsRejected()
	getsDropped := c.cache.Metrics.GetsDropped()
	setsDropped := c.cache.Metrics.SetsDropped()
	ratio := c.cache.Metrics.Ratio()

	return cache2.CacheStats{
		HitCount:    int64(hits),
		MissCount:   int64(misses),
		KeysAdded:   int64(keyAdd),
		KeysUpdate:  int64(updated),
		KeysEvict:   int64(keysEvicted),
		CostAdd:     int64(costAdd),
		CostEvict:   int64(costEvicted),
		RejectSets:  int64(rejected),
		GetDropGets: int64(getsDropped),
		SetDropGets: int64(setsDropped),
		KeepGets:    int64(kept),
		Ratio:       ratio,
	}
}
*/
