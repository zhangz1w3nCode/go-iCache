package imp

//import (
//	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
//	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
//	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-factory"
//	"sync"
//)
//
//type AdapterCacheFactory struct {
//	factories map[string]cache-factory.CacheFactory
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
//func (a *AdapterCacheFactory) GetCache(cache-config *cache-config.GoCacheConfig) iCache.ICache {
//	for _, factoryItem := range a.factories {
//		if factoryItem.Support(cache-config.CacheType) {
//			return factoryItem.GetCache(cache-config)
//		}
//	}
//	return nil
//}
//
//// RegisterFactory 注册缓存工厂
//func (a *AdapterCacheFactory) RegisterFactory(cache-factory cache-factory.CacheFactory) {
//	for _, f := range a.factories {
//		if f.Support(cache-factory.Support("")) {
//			panic("Duplicate cache cache-factory registered")
//		}
//	}
//	a.factories["go_cache_factory"] = cache-factory
//}
//
//// GetInstance 获取缓存工厂适配器实例
//func (a *AdapterCacheFactory) GetInstance() *AdapterCacheFactory {
//	return adapter
//}
