package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/register"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	"time"
)

type TestService struct {
	helloworld.UnimplementedTestServiceServer
}

func NewTestService() *TestService {
	return &TestService{}
}

func (s *TestService) CreateCache(ctx context.Context,
	in *helloworld.CreateCacheRequest) (*helloworld.CreateCacheReply, error) {

	cacheConfig := config.GoCacheConfig{
		CacheName:  in.CacheName,
		CacheType:  "go_cache",
		ExpireTime: 5 * time.Minute,
		CleanTime:  5 * time.Minute,
	}

	cacheAPI := register.CacheAPI
	cache, err := cacheAPI.CreateCache(cacheConfig)

	if err != nil {
		return &helloworld.CreateCacheReply{Message: "创建缓存失败"}, nil
	}

	if cache == nil {
		return &helloworld.CreateCacheReply{Message: "创建缓存失败"}, nil
	} else {
		return &helloworld.CreateCacheReply{Message: "创建缓存成功"}, nil
	}
}

func (s *TestService) GetCacheKey(ctx context.Context,
	in *helloworld.GetCacheKeyRequest) (*helloworld.GetCacheKeyReply, error) {

	cacheAPI := register.CacheAPI

	value, err := cacheAPI.GetCacheKey(in.GetCacheKey(), in.GetCacheName())

	if err != nil {
		return &helloworld.GetCacheKeyReply{CacheValue: "cache is empty"}, nil
	}

	if value == nil {
		return &helloworld.GetCacheKeyReply{CacheValue: "cache value is empty"}, nil
	}

	return &helloworld.GetCacheKeyReply{CacheValue: value.GetCacheValue()}, nil
}

func (s *TestService) SetCacheKey(ctx context.Context,
	in *helloworld.SetCacheKeyRequest) (*helloworld.SetCacheKeyReply, error) {

	cacheAPI := register.CacheAPI

	resp, err := cacheAPI.SetCacheKey(in.GetCacheName(), in.GetCacheKey(), in.GetCacheVal())

	if err != nil {
		return &helloworld.SetCacheKeyReply{CacheVa: "-1"}, nil
	}

	return &helloworld.SetCacheKeyReply{CacheVa: resp.GetCacheVa()}, nil
}
