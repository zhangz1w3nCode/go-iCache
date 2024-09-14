package factory

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
)

// CacheFactory 缓存工厂接口
type CacheFactory interface {
	Support(cacheType string) bool
	GetCache(config *config.GoCacheConfig) iCache.ICache
}
