package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	"time"
)

type TestService struct {
	helloworld.UnimplementedTestServiceServer
	manager *manager.CacheManager
}

func NewTestService(manager *manager.CacheManager) *TestService {
	return &TestService{
		manager: manager,
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

	cache := s.manager.CreateCache(cacheConfig)

	if cache == nil {
		return &helloworld.CreateCacheReply{Message: "创建缓存失败"}, nil
	} else {
		return &helloworld.CreateCacheReply{Message: "创建缓存成功"}, nil
	}
}

func (s *TestService) GetCacheKey(ctx context.Context,
	in *helloworld.GetCacheKeyRequest) (*helloworld.GetCacheKeyReply, error) {

	cache := s.manager.GetCache(in.GetCacheName())

	if cache == nil {
		return &helloworld.GetCacheKeyReply{CacheValue: "cache value is empty"}, nil
	}

	value := cache.Get(in.GetCacheKey())

	return &helloworld.GetCacheKeyReply{CacheValue: value.Data.(string)}, nil
}

func (s *TestService) SetCacheKey(ctx context.Context,
	in *helloworld.SetCacheKeyRequest) (*helloworld.SetCacheKeyReply, error) {

	cache := s.manager.GetCache(in.GetCacheName())

	if cache == nil {
		return &helloworld.SetCacheKeyReply{CacheVa: "cache value is empty"}, nil
	}

	cache.Set(in.GetCacheKey(), in.GetCacheVal())

	return &helloworld.SetCacheKeyReply{CacheVa: "set successfully"}, nil
}
