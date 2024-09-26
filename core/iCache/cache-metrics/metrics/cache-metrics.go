package metrics

import (
	"errors"
	"sync"
)

// CacheMetrics 整个缓存的状态
type CacheMetrics struct {
	mu                   sync.Mutex
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

// DetectCacheKeyCount 检测key metric是否超过阈值
func (c *CacheMetrics) DetectCacheKeyCount(preNode *CacheMetrics, limit int64, thresholdRate float64, thresholdRateHistory float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 基于数量
	if c.CacheCurrentKeyCount > limit { // 假设阈值为10000
		return errors.New("cache is more than CacheMaxCount")
	}

	// 基于百分比
	if float64(c.CacheCurrentKeyCount)/float64(c.CacheMaxCount) > thresholdRate {
		return errors.New("cache is more than 80% full")
	}

	// 基于前一个收集到的指标的判断
	previousCount := preNode.CacheCurrentKeyCount
	if previousCount == 0 {
		return errors.New("previous cache size is 0")
	}
	if (float64(c.CacheCurrentKeyCount-previousCount) / float64(previousCount)) > thresholdRateHistory {
		return errors.New("cache size has increased by more than thresholdRateHistory")
	}

	return nil
}

// DetectCacheQueryCount 检测query metric是否超过阈值
func (c *CacheMetrics) DetectCacheQueryCount(preNode *CacheMetrics, limit int64, thresholdRate float64, thresholdRateHistory float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 基于数量
	if c.CacheHitCount > limit { // 假设阈值为10000
		return errors.New("cache is more than CacheMaxCount")
	}

	// 基于百分比
	if float64(c.CacheHitCount)/float64(c.CacheQueryCount) > thresholdRate {
		return errors.New("cache is more than 80% full")
	}

	// 基于前一个收集到的指标的判断
	previousHitCount := preNode.CacheHitCount
	if (float64(c.CacheHitCount-previousHitCount) / float64(previousHitCount)) > thresholdRateHistory {
		return errors.New("cache size has increased by more than 70%")
	}

	return nil
}

// additional check
