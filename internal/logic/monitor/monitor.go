package test

import (
	"encoding/json"
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MonitorLogic struct {
	monitor monitorpb.UnimplementedCacheMonitorServiceServer
	manager *cacheManager.CacheManager
}

func NewMonitorService(manager *cacheManager.CacheManager) *MonitorLogic {
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
		return nil, status.Errorf(codes.Unavailable, "Get cache "+cacheName+" from cache cache-manager error!")
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
		return "", status.Errorf(codes.Unavailable, "Get cache "+cacheName+" from cache cache-manager error!")
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

// GetCacheMetrics 	获取缓存的指标
func (m *MonitorLogic) GetCacheMetrics(cacheName string) (*monitorpb.CacheMetrics, error) {
	cache := m.manager.GetCache(cacheName)
	if cache == nil {
		return nil, status.Errorf(codes.Unavailable, "Get cache "+cacheName+" from cache cache-manager error!")
	}
	cacheMetrics := cache.GetCacheMetrics()
	if cacheMetrics == nil {
		return nil, nil
	}

	metrics := &monitorpb.CacheMetrics{
		CacheName:            cacheName,
		CacheSize:            cacheMetrics.CacheSize,
		CacheHitCount:        cacheMetrics.CacheHitCount,
		CacheMissCount:       cacheMetrics.CacheMissCount,
		CacheQueryCount:      cacheMetrics.CacheQueryCount,
		CacheCurrentKeyCount: cacheMetrics.CacheCurrentKeyCount,
		CacheMaxCount:        cacheMetrics.CacheMaxCount,
		CacheHitRate:         cacheMetrics.CacheHitRate,
		CacheMissRate:        cacheMetrics.CacheMissRate,
	}

	return metrics, nil
}
