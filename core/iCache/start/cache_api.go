package start

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
)

type CacheApi struct {
	init *CacheInit
}

func NewCacheApi(mm *manager.CacheManager) *CacheApi {
	return &CacheApi{
		init: &CacheInit{
			CacheManager: mm,
		},
	}
}

func (s *CacheApi) CreateCache(config config.GoCacheConfig) (*helloworld.CreateCacheReply, error) {
	createCache := s.init.CacheManager.CreateCache(config)
	if createCache == nil {
		return &helloworld.CreateCacheReply{Message: "创建缓存失败"}, nil
	} else {
		return &helloworld.CreateCacheReply{Message: "创建缓存成功"}, nil
	}
}

func (s *CacheApi) GetCacheKey(cacheName string, cacheKey string) (*helloworld.GetCacheKeyReply, error) {
	cache := s.init.CacheManager.GetCache(cacheName)

	if cache == nil {
		return &helloworld.GetCacheKeyReply{CacheValue: "cache is empty"}, nil
	}

	value := cache.Get(cacheKey)
	if value == nil {
		return &helloworld.GetCacheKeyReply{CacheValue: "cache value is empty"}, nil
	}

	return &helloworld.GetCacheKeyReply{CacheValue: value.Data.(string)}, nil
}

func (s *CacheApi) SetCacheKey(cacheName string, cacheKey string, cacheValue interface{}) (*helloworld.SetCacheKeyReply, error) {

	cache := s.init.CacheManager.GetCache(cacheName)

	if cache == nil {
		return &helloworld.SetCacheKeyReply{CacheVa: "-1"}, nil
	}

	cache.Set(cacheKey, cacheValue)

	return &helloworld.SetCacheKeyReply{CacheVa: "successfully"}, nil
}
