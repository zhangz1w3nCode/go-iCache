package imp

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/constant"
	go_cache "github.com/zhangz1w3nCode/go-iCache/core/iCache/go-cache"
)

// GoCacheFactory 是 Caffeine 缓存工厂实现
type GoCacheFactory struct{}

func (f *GoCacheFactory) Support(cacheType string) bool {
	return cacheType == constant.CACHE_TYPE_GO_CACHE
}

func (f *GoCacheFactory) GetCache(config *config.GoCacheConfig) iCache.ICache {
	return go_cache.NewGoCache(config)
}
