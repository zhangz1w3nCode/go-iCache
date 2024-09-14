package init

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
)

type CacheInit struct {
	CacheManager *manager.CacheManager
}

func NewCacheInit() *CacheInit {
	return &CacheInit{
		CacheManager: manager.NewCacheManager(),
	}
}

func (c *CacheInit) CreatCache(config config.GoCacheConfig) *iCache.ICache {
	cache := c.CacheManager.CreateCache(config)
	return &cache
}
