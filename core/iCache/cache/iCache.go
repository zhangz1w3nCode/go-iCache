package cache

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/value-wrapper"
)

// ICache 真正缓存的接口
type ICache interface {
	Get(key string) *value_wrapper.ValueWrapper
	Set(key string, value interface{})
	GetValues() []*value_wrapper.ValueWrapper
	GetKeys() []string
	Size() int
	GetName() string
	CalculateMemoryUsage() float64
	GetCacheStatus() CacheStats
}

// CacheStats 缓存状态统计
type CacheStats struct {
	HitCount   int64
	MissCount  int64
	KeysAdded  int64
	KeysUpdate int64
	KeysEvict  int64
	CostAdd    int64
	CostEvict  int64
	RejectSets int64
	// The following 2 keep track of how many gets were kept and dropped on the
	// floor.
	GetDropGets int64
	SetDropGets int64
	KeepGets    int64
	// 命中率
	Ratio float64
}
