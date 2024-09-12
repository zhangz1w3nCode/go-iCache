package iCache

// RealCache 真正缓存的接口
type RealCache interface {
	Get(key string) *ValueWrapper
	Set(key string, value interface{})
	GetValues() []*ValueWrapper
	GetKeys() []string
	Size() int
	GetName() string
	CalculateMemoryUsage() float64
	GetCacheStatus() CacheStats
}

// CacheStats 缓存状态统计
type CacheStats struct {
	HitCount  int64
	MissCount int64
}
