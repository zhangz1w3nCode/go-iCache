package cacheFactory

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache"
	cacheConfig "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
)

// CacheFactory 缓存工厂接口
type CacheFactory interface {
	Support(cacheType string) bool
	GetCache(config *cacheConfig.GoCacheConfig) cache.ICache
}
