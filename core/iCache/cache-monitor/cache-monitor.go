package monitor

import (
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/collector"
	goCache "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/go-cache"
	"log"
	"sync"
	"time"
)

type CacheMonitor struct {
	frequency      time.Duration
	manager        *cacheManager.CacheManager
	cacheName      string
	cacheCollector *collector.MetricCollector
	ticker         *time.Ticker
	wg             *sync.WaitGroup
}

func NewCacheMonitor(frequency time.Duration, manager *cacheManager.CacheManager, cacheName string) *CacheMonitor {
	return &CacheMonitor{
		frequency:      frequency,
		manager:        manager,
		cacheName:      cacheName,
		cacheCollector: collector.NewMetricCollector(manager, frequency, 10),
		ticker:         time.NewTicker(frequency),
	}
}

func (c *CacheMonitor) Start() error {
	//初始化
	c.cacheCollector.InitCollection(c.cacheName)
	defer c.ticker.Stop()

	// 使用协程异步运行定时任务
	//c.wg.Add(1)
	go func() {
		//defer c.wg.Done()
		for range c.ticker.C {
			// 获取缓存管理器的所有缓存
			cacheDetail := c.manager.GetCacheDetail()
			if cacheDetail == nil {
				log.Fatalf("cache detail is nil")
			}
			if cacheDetail[c.cacheName] == nil {
				log.Fatalf("cache is no exist! cacheName: [%s]", c.cacheName)
			}
			cacheInstance := cacheDetail[c.cacheName].(*goCache.GoCache)
			// 调用每个缓存的监控方法得到监控指标的来源
			metric := cacheInstance.GetCacheMetrics()
			// 进入指标采集
			_ = c.cacheCollector.CollectCacheKeyCount(metric, nil, 1000, 0.5, 0.3)
			// _ = c.cacheCollector.CollectCacheHitCount(metric, nil, 10000, 0.8, 0.25)
		}
	}()

	// 等待协程完成
	//c.wg.Wait()
	return nil
}
