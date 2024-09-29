package collector

import (
	"errors"
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	cacheMcs "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/metrics"
	goCache "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/go-cache"
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

// Collect 收集指标
func (c *MetricCollector) Collect(metric *cacheMcs.CacheMetrics, wg *sync.WaitGroup) error {
	//get pre node
	if len(c.History) <= 0 {
		log.Fatalf("pre node is nil")
	}
	//at lease one
	preNode := c.History[len(c.History)-1]
	//metrics rule match
	CacheKeyCountErr := preNode.DetectCacheKeyCount(preNode, 80, 0.8, 0.25)
	CacheQueryCountErr := preNode.DetectCacheQueryCount(preNode, 80, 0.8, 0.25)

	if CacheKeyCountErr != nil {
		log.Printf("detect cache key count metrics error: %s", CacheKeyCountErr.Error())
	}
	if CacheQueryCountErr != nil {
		log.Printf("detect cache query count metrics error: %s", CacheQueryCountErr.Error())
	}

	//add history
	c.addToHistoryArray(metric)

	return nil
}

// CollectCacheKeyCount 收集指标
func (c *MetricCollector) CollectCacheKeyCount(metric *cacheMcs.CacheMetrics, limit int64, thresholdRate float64, thresholdRateHistory float64) error {
	//get pre node
	if len(c.History) <= 0 {
		log.Fatalf("pre node is nil")
	}
	//at lease one
	preNode := c.History[len(c.History)-1]
	//detect cache hit count
	err := detect(preNode.CacheCurrentKeyCount, metric.CacheCurrentKeyCount, limit, metric.CacheMaxCount, thresholdRate, thresholdRateHistory)
	if err != nil {
		log.Printf("detect cache key count metrics error: %s", err.Error())
		return nil
	}
	//add history array
	c.addToHistoryArray(metric)

	return nil
}

// CollectCacheHitCount 收集指标
func (c *MetricCollector) CollectCacheHitCount(metric *cacheMcs.CacheMetrics, wg *sync.WaitGroup, limit int64, thresholdRate float64, thresholdRateHistory float64) error {
	//get pre node
	if len(c.History) <= 0 {
		log.Fatalf("pre node is nil")
	}
	//at lease one
	currentNodeCount := len(c.History) - 1

	//v1 参考前一个节点的值
	preNode := c.History[len(c.History)-1]

	//v2 参考前n个节点的某个指标的平均值
	avgMetric := int64(0)
	sumMetric := int64(0)
	for i := 0; i < currentNodeCount; i++ {
		sumMetric += preNode.CacheHitCount
	}
	avgMetric = sumMetric / int64(currentNodeCount)

	//探测 cache hit count
	err := detect(avgMetric, metric.CacheHitCount, limit, metric.CacheQueryCount, thresholdRate, thresholdRateHistory)
	if err != nil {
		log.Printf("detect cache hit count metrics error: %s", err.Error())
		return nil
	}
	//add history
	c.addToHistoryArray(metric)

	return nil
}

func (c *MetricCollector) InitCollection(cacheName string) {
	//初始化：每3秒 采集一次指标 采集10次 作为初始化的队列 填满历史数组
	for i := 1; i <= 10; i++ {
		//获取缓存管理器的所有缓存
		cacheDetail := c.manager.GetCacheDetail()
		if cacheDetail == nil {
			log.Fatalf("cache detail is nil")
		}
		if cacheDetail[cacheName] == nil {
			log.Fatalf("cache is no exist!cacheName:[%s]", cacheName)
		}
		cacheInstance := cacheDetail[cacheName].(*goCache.GoCache)
		//调用每个缓存的监控方法得到监控指标的来源
		metric := cacheInstance.GetCacheMetrics()
		//采集
		c.addToHistoryArray(metric)
		//模拟定时任务 3秒入队一次
		time.Sleep(3 * time.Second)
	}
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

// Detect 检测 metric是否超过阈值
func detect(preVal int64, current int64, limit int64, all int64, thresholdRate float64, thresholdRateHistory float64) error {
	// 基于数量
	if current > limit { // 假设阈值为10000
		return errors.New("monitor metric current value more than limit")
	}

	// 基于百分比
	if float64(current)/float64(all) > thresholdRate {
		return errors.New("monitor metric current value owner rate more than thresholdRate")
	}

	// 比之前的指标 增长或下降 超过阈值
	if (float64(current-preVal) / float64(preVal)) >= thresholdRateHistory {
		return errors.New("monitor metric current value has increment more than thresholdRateHistory")
	}

	// 比之前的指标 增长或下降 超过阈值
	if (float64(current+preVal) / float64(preVal)) >= thresholdRateHistory {
		return errors.New("monitor metric current value has  decrement more than thresholdRateHistory")
	}

	return nil
}
