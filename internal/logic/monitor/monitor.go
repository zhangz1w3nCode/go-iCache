package test

import (
	"encoding/json"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MonitorLogic struct {
	monitor monitorpb.UnimplementedCacheMonitorServiceServer
	manager *manager.CacheManager
}

func NewMonitorService(manager *manager.CacheManager) *MonitorLogic {
	return &MonitorLogic{
		manager: manager,
	}
}

// GetCacheNameList 获取某个appName下的某台机器上有哪些缓存：如productCache、userCache
func (m *MonitorLogic) GetCacheNameList() ([]string, error) {
	cacheName := m.manager.GetAllCacheName()
	if cacheName == nil || len(cacheName) == 0 {
		return []string{}, nil
	}
	return cacheName, nil
}

// GetCacheKeyList 获取cache的key列表
func (m *MonitorLogic) GetCacheKeyList(cacheName string) ([]string, error) {
	cache := m.manager.GetCache(cacheName)
	if cache == nil {
		return nil, status.Errorf(codes.Unavailable, "Get cache "+cacheName+" from cache manager error!")
	}
	keyList := cache.GetKeys()
	if keyList == nil || len(keyList) == 0 {
		return []string{}, nil
	}
	return keyList, nil
}

// GetValueToCacheUser 获取某个缓存的value
func (m *MonitorLogic) GetValueToCacheUser(cacheName string, cacheKey string) (string, error) {
	cache := m.manager.GetCache(cacheName)
	if cache == nil {
		return "", status.Errorf(codes.Unavailable, "Get cache "+cacheName+" from cache manager error!")
	}
	valueWrapper := cache.Get(cacheKey)
	if valueWrapper == nil {
		return "{}", nil
	}
	valueByte, err := json.Marshal(valueWrapper)
	if err != nil {
		return "{}", err
	}
	valueStr := string(valueByte)
	if valueStr == "" {
		return "{}", nil
	}
	return valueStr, nil
}
