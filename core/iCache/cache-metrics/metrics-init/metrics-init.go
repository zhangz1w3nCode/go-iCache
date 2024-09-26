package metricsInit

import (
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/collector"
	goCache "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/go-cache"
	"log"
	"time"
)





func CreateCacheMonitor(frequency time.Duration, manager *cacheManager.CacheManager, cacheName string) error {

	productCacheMetricsCollector := collector.NewMetricCollector(manager, frequency, 10)

	productCacheMetricsCollector.InitCollection(cacheName)

	//通过定时任务，定时收集指标
	ticker := time.NewTicker(frequency)
	defer ticker.Stop()

	for _ = range ticker.C {
		//获取缓存管理器的所有缓存
		cacheDetail := manager.GetCacheDetail()

		if cacheDetail == nil {
			log.Fatalf("cache detail is nil")
		}
		if cacheDetail[cacheName] == nil {
			log.Fatalf("cache is no exist!cacheName:[%s]", cacheName)
		}
		cacheInstance := cacheDetail[cacheName].(*goCache.GoCache)
		//调用每个缓存的监控方法得到监控指标的来源
		metric := cacheInstance.GetCacheMetrics()
		//进入指标采集
		_ = productCacheMetricsCollector.Collect(metric, nil)

	}

	return nil
}
