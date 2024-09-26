package cache

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/metrics"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/value-wrapper"
)

// ICache 真正缓存的接口
type ICache interface {
	Get(key string) *valueWrapper.ValueWrapper
	Set(key string, value interface{})
	GetValues() []*valueWrapper.ValueWrapper
	GetKeys() []string
	CacheNum() int
	GetName() string
	GetCacheValuesStatus() []*metrics.CacheValueMetrics
	GetCacheMetrics() *metrics.CacheMetrics
}
