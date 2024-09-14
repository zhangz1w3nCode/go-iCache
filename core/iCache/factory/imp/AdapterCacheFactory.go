package imp

//import (
//	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
//	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
//	"github.com/zhangz1w3nCode/go-iCache/core/iCache/factory"
//	"sync"
//)
//
//type AdapterCacheFactory struct {
//	factories map[string]factory.CacheFactory
//	lock      sync.RWMutex
//}
//
//func (a *AdapterCacheFactory) Support(cacheType string) bool {
//	for _, factoryItem := range a.factories {
//		if factoryItem.Support(cacheType) {
//			return true
//		}
//	}
//	return false
//}
//
//func (a *AdapterCacheFactory) GetCache(config *config.GoCacheConfig) iCache.ICache {
//	for _, factoryItem := range a.factories {
//		if factoryItem.Support(config.CacheType) {
//			return factoryItem.GetCache(config)
//		}
//	}
//	return nil
//}
//
//// RegisterFactory 注册缓存工厂
//func (a *AdapterCacheFactory) RegisterFactory(factory factory.CacheFactory) {
//	for _, f := range a.factories {
//		if f.Support(factory.Support("")) {
//			panic("Duplicate cache factory registered")
//		}
//	}
//	a.factories["go_cache_factory"] = factory
//}
//
//// GetInstance 获取缓存工厂适配器实例
//func (a *AdapterCacheFactory) GetInstance() *AdapterCacheFactory {
//	return adapter
//}
