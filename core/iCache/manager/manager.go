package manager

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/factory"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/factory/imp"
)

// CacheManager 缓存管理器
type CacheManager struct {
	cacheMap  map[string]iCache.ICache
	factory   factory.CacheFactory
	configMap map[string]config.GoCacheConfig
}

// NewCacheManager 创建一个新的 CacheManager 实例
func NewCacheManager() *CacheManager {
	return &CacheManager{
		cacheMap:  make(map[string]iCache.ICache),
		factory:   &imp.GoCacheFactory{},
		configMap: make(map[string]config.GoCacheConfig),
	}
}

func (m *CacheManager) CreateCache(config config.GoCacheConfig) iCache.ICache {
	value, exists := m.cacheMap[config.CacheName]
	if exists {
		return value
	}
	cacheInstance := m.factory.GetCache(&config)
	m.cacheMap[config.CacheName] = cacheInstance
	m.configMap[config.CacheName] = config
	return cacheInstance
}

func (m *CacheManager) GetCache(cacheName string) iCache.ICache {
	return m.cacheMap[cacheName]
}
