package collector

import (
	"errors"
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	cacheMcs "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/metrics"
	"log"
	"sync"
	"time"
)

// MetricCollector 指标收集器
type MetricCollector struct {
	mu           sync.Mutex
	History      []*cacheMcs.CacheMetrics
	capacity     int64
	currentCount int64
	//additional files
	cycle   time.Duration
	manager *cacheManager.CacheManager
}

// NewMetricCollector 新建指标收集器
func NewMetricCollector(manager *cacheManager.CacheManager, cycle time.Duration, capacity int64) *MetricCollector {
	return &MetricCollector{
		History:      make([]*cacheMcs.CacheMetrics, 0, capacity),
		capacity:     capacity,
		cycle:        cycle,
		currentCount: 0,
		manager:      manager,
	}
}

// CollectCacheKeyCount 收集cache key count指标
func (c *MetricCollector) CollectCacheKeyCount(metric *cacheMcs.CacheMetrics,
	limit float64, thresholdRate float64,
	thresholdRateHistory float64, changer cacheMcs.CacheMetricsChanger) error {

	currentArrayCount := len(c.History)
	if int64(currentArrayCount) == c.capacity {
		currentNodeCount := currentArrayCount
		// 参考前n个节点的某个指标的平均值
		avgMetric := float64(0)
		sumMetric := float64(0)
		//todo: 这里可以优化 使用前缀和 不需要每次都重复计算
		for _, value := range c.History {
			sumMetric += changer(value)
		}
		avgMetric = sumMetric / float64(currentNodeCount)
		err := c.MetricsDetector(avgMetric, float64(metric.CacheCurrentKeyCount), limit, float64(metric.CacheMaxCount), thresholdRate, thresholdRateHistory)
		if err != nil {
			log.Printf("detect cache key count metrics error: %s", err.Error())
			return err
		}
	}
	c.addToHistoryArray(metric)
	return nil
}

func (c *MetricCollector) addToHistoryArray(metrics *cacheMcs.CacheMetrics) {
	//add history
	c.History = append(c.History, metrics)
	c.currentCount++
	if int64(len(c.History)) > c.capacity {
		c.History = c.History[1:]
		c.currentCount = int64(len(c.History))
	}
}

// MetricsDetector 指标探测器 检测 metric是否超过阈值
func (c *MetricCollector) MetricsDetector(avg float64, current float64, limit float64, all float64, thresholdRate float64, thresholdRateHistory float64) error {
	// 基于数量
	if current > limit { // 假设阈值为10000
		return errors.New("monitor metric current value more than limit")
	}

	// 基于百分比
	if current/all > thresholdRate {
		return errors.New("monitor metric current value owner rate more than thresholdRate")
	}

	// 比之前的指标 增长 超过阈值
	if (current-avg) != 0 && ((current-avg)/avg) >= thresholdRateHistory {
		return errors.New("monitor metric current value has increment more than thresholdRateHistory")
	}

	// 比之前的指标 下降 超过阈值
	if (current-avg) != 0 && ((avg-current)/avg) >= thresholdRateHistory {
		return errors.New("monitor metric current value has decrement more than thresholdRateHistory")
	}

	return nil
}
