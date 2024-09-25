package cacheMetrics

// CacheValueMetrics 整个缓存的状态
type CacheValueMetrics struct {
	CacheSize     float32 //当前这个key的大小
	CacheHitCount int64
}

func NewCacheValueMetrics() *CacheValueMetrics {
	return &CacheValueMetrics{
		CacheSize:     0,
		CacheHitCount: 0,
	}
}
