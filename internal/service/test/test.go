package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	"github.com/zhangz1w3nCode/go-iCache/internal/logic/cache"
	"time"
)

type TestService struct {
	helloworld.UnimplementedTestServiceServer
	logic   *cache.TestLogic
	manager *manager.CacheManager
}

func NewTestService(mm *manager.CacheManager) *TestService {
	return &TestService{
		logic:   cache.NewTestLogic(),
		manager: mm,
	}
}

func (s *TestService) CreateCache(ctx context.Context,
	in *helloworld.CreateCacheRequest) (*helloworld.CreateCacheReply, error) {
	cacheConfig := config.GoCacheConfig{
		CacheName:  in.CacheName,
		CacheType:  "go_cache",
		ExpireTime: 5 * time.Minute,
		CleanTime:  5 * time.Minute,
	}
	createCache := s.manager.CreateCache(cacheConfig)

	if createCache == nil {
		return &helloworld.CreateCacheReply{Message: "创建缓存失败"}, nil
	} else {
		return &helloworld.CreateCacheReply{Message: "创建缓存成功"}, nil
	}
}

func (s *TestService) GetCacheKey(ctx context.Context,
	in *helloworld.GetCacheKeyRequest) (*helloworld.GetCacheKeyReply, error) {
	cache := s.manager.GetCache(in.GetCacheName())

	if cache == nil {
		return &helloworld.GetCacheKeyReply{CacheValue: "-1"}, nil
	}

	value := cache.Get(in.GetCacheKey())
	if value == nil {
		return &helloworld.GetCacheKeyReply{CacheValue: "cache value is empty"}, nil
	}

	return &helloworld.GetCacheKeyReply{CacheValue: value.Data.(string)}, nil
}

func (s *TestService) SetCacheKey(ctx context.Context,
	in *helloworld.SetCacheKeyRequest) (*helloworld.SetCacheKeyReply, error) {

	cache := s.manager.GetCache(in.GetCacheName())

	if cache == nil {
		return &helloworld.SetCacheKeyReply{CacheVa: "-1"}, nil
	}

	cache.Set(in.GetCacheKey(), in.GetCacheVal())

	return &helloworld.SetCacheKeyReply{CacheVa: "successfully"}, nil
}
