package metrics

// CacheMetrics 整个缓存的状态
type CacheMetrics struct {
	CacheMaxCount        int64
	CacheCurrentKeyCount int64
	CacheSize            int64
	CacheHitCount        int64
	CacheHitRate         float32
	CacheMissCount       int64
	CacheMissRate        float32
	CacheQueryCount      int64
}

func NewCacheMetrics(cacheMaxCount int64) *CacheMetrics {
	return &CacheMetrics{
		CacheMaxCount:        cacheMaxCount,
		CacheCurrentKeyCount: 0,
		CacheSize:            0,
		CacheHitCount:        0,
		CacheMissCount:       0,
		CacheQueryCount:      0,
	}
}

// CacheMetricsChanger 指标选择器：选择计算哪个指标
type CacheMetricsChanger func(interface{}) float64
