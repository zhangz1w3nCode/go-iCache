package imp

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache"
	cacheConfig "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
	cacheConstant "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-constant"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/go-cache"
)

// GoCacheFactory 是 GoCache 缓存工厂实现
type GoCacheFactory struct{}

func (f *GoCacheFactory) Support(cacheType string) bool {
	return cacheType == cacheConstant.CACHE_TYPE_GO_CACHE
}

func (f *GoCacheFactory) GetCache(config *cacheConfig.GoCacheConfig) cache.ICache {
	return goCache.NewGoCache(config)
}
