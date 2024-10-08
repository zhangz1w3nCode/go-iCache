package collector

import (
	"errors"
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	"log"
	"math"
	"sync"
	"time"
)

// MetricCollector 指标收集器
type MetricCollector struct {
	mu                 sync.Mutex
	metricHistoryQueue []float64
	capacity           int64
	currentCount       int64
	//additional files
	cycle   time.Duration
	manager *cacheManager.CacheManager
}

// NewMetricCollector 新建指标收集器
func NewMetricCollector(manager *cacheManager.CacheManager, cycle time.Duration, capacity int64) *MetricCollector {
	return &MetricCollector{
		metricHistoryQueue: make([]float64, 0, capacity),
		capacity:           capacity,
		cycle:              cycle,
		currentCount:       0,
		manager:            manager,
	}
}

// CollectCacheQueryCount 收集cache key query count指标
func (c *MetricCollector) CollectCacheQueryCount(currentMetric float64,
	max float64, thresholdRate float64,
	thresholdRateHistory float64) error {

	defer func() {
		c.addToMetricHistoryQueue(currentMetric)
	}()

	currentArrayCount := len(c.metricHistoryQueue)

	if int64(currentArrayCount) == c.capacity {
		currentNodeCount := currentArrayCount
		// 参考前n个节点的某个指标的平均值
		avgMetric := float64(0)
		sumMetric := float64(0)
		//todo: 这里可以优化 使用前缀和 不需要每次都重复计算
		for _, value := range c.metricHistoryQueue {
			sumMetric += value
		}
		avgMetric = sumMetric / float64(currentNodeCount)
		err := c.MetricsDetector(avgMetric, currentMetric, max, thresholdRate, thresholdRateHistory)
		if err != nil {
			log.Printf("detect cache query count metrics error: %s", err.Error())
			return err
		}
	}

	return nil
}

func (c *MetricCollector) addToMetricHistoryQueue(currentMetric float64) {
	c.metricHistoryQueue = append(c.metricHistoryQueue, currentMetric)
	c.currentCount++
	if int64(len(c.metricHistoryQueue)) > c.capacity {
		c.metricHistoryQueue = c.metricHistoryQueue[1:]
		c.currentCount = int64(len(c.metricHistoryQueue))
	}
}

// MetricsDetector 指标探测器 检测 metric是否超过阈值
func (c *MetricCollector) MetricsDetector(avg float64, current float64, max float64, thresholdRate float64, thresholdRateHistory float64) error {
	if current == 0 {
		//log.Printf("monitor metric current value is 0")
	}

	// 基于数量
	if current > max { // 假设阈值为10000
		return errors.New("monitor metric current value more than limit")
	}

	// 基于百分比
	if current/max > thresholdRate {
		return errors.New("monitor metric current value owner rate more than thresholdRate")
	}

	//没有增长或者下降 直接结束
	if current == avg {
		if current == float64(0) || avg == float64(0) {
			//为0 没有变动
			log.Printf("monitor current metric value equals avg! current:%f,avg:%f", current, avg)
		} else {
			return nil
		}

	}
	//对比之前平均值有增长 同时 增长率超过阈值
	diff := math.Abs(current - avg)
	if (diff / avg) > thresholdRateHistory {
		log.Printf("current value is %f,history avg value is %f,increment or decrement rate is %f,thresholdRateHistory is %f", current, avg, diff/avg, thresholdRateHistory)
		//return errors.New("monitor metric current value has increment or decrement more than thresholdRateHistory")
	}

	return nil
}
