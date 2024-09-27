package monitor

import (
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/collector"
	goCache "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/go-cache"
	"log"
	"time"
)

type CacheMonitor struct {
	frequency      time.Duration
	manager        *cacheManager.CacheManager
	cacheName      string
	cacheCollector *collector.MetricCollector
	ticker         *time.Ticker
}

func NewCacheMonitor(frequency time.Duration, manager *cacheManager.CacheManager, cacheName string) *CacheMonitor {
	return &CacheMonitor{
		frequency:      frequency,
		manager:        manager,
		cacheName:      cacheName,
		cacheCollector: collector.NewMetricCollector(manager, frequency, 10),
	}
}

func (c *CacheMonitor) Start() {
	//初始化
	c.cacheCollector.InitCollection(c.cacheName)

	// 启动定时器
	c.ticker = time.NewTicker(c.frequency)
	defer c.ticker.Stop()

	func() {
		// 使用协程异步运行定时任务
		for { //无限循环
			<-c.ticker.C    //这里定义了一个case，监听名为ticker的定时器生成的channel
			c.MonitorTask() //执行定时任务
		}
	}()
}

func (c *CacheMonitor) MonitorTask() {
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
	_ = c.cacheCollector.CollectCacheKeyCount(metric, 1000, 0.5, 0.3)
	// _ = c.cacheCollector.CollectCacheHitCount(metric, nil, 10000, 0.8, 0.25)
}
