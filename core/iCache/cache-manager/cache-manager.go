package cacheManager

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache"
	cacheConfig "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
	cacheFactory "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-factory"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-factory/imp"
)

// CacheManager 缓存管理器
type CacheManager struct {
	cacheMap  map[string]cache.ICache
	factory   cacheFactory.CacheFactory
	configMap map[string]cacheConfig.GoCacheConfig
}

// NewCacheManager 创建一个新的 CacheManager 实例
func NewCacheManager() *CacheManager {
	return &CacheManager{
		cacheMap:  make(map[string]cache.ICache),
		factory:   &imp.GoCacheFactory{},
		configMap: make(map[string]cacheConfig.GoCacheConfig),
	}
}

func (m *CacheManager) CreateCache(config cacheConfig.GoCacheConfig) cache.ICache {
	value, exists := m.cacheMap[config.CacheName]
	if exists {
		return value
	}
	cacheInstance := m.factory.GetCache(&config)
	m.cacheMap[config.CacheName] = cacheInstance
	m.configMap[config.CacheName] = config
	return cacheInstance
}

func (m *CacheManager) GetCache(cacheName string) cache.ICache {
	cache, exists := m.cacheMap[cacheName]
	if !exists {
		return nil
	}
	return cache
}

func (m *CacheManager) GetAllCacheName() []string {
	var cacheNameList []string
	for key := range m.cacheMap {
		cacheNameList = append(cacheNameList, key)
	}
	return cacheNameList
}
func (m *CacheManager) GetCacheDetail() map[string]cache.ICache {
	return m.cacheMap
}
