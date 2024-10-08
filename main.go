package main

import (
	"fmt"
	cacheConfig "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-config"
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	monitor "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-monitor"
	"time"
)

func main() {
	//mock cache user
	manager := cacheManager.NewCacheManager()
	cache := manager.CreateCache(cacheConfig.GoCacheConfig{
		CacheName:     "user-cache",
		CacheType:     "go-cache",
		CacheMaxCount: 1000,
		ExpireTime:    20 * time.Second,
		CleanTime:     20 * time.Second,
	})

	for i := 0; i < 500; i++ {
		cache.Set(fmt.Sprintf("user-%d", i), i)
	}
	for i := 0; i < 900; i++ {
		cache.Get(fmt.Sprintf("user-%d", i))
	}

	userCacheMonitor := monitor.NewCacheMonitor(1*time.Second, manager, "user-cache")

	go userCacheMonitor.Start()

	for k := 0; true; k++ {
		fmt.Println(k)
		time.Sleep(1 * time.Second)
	}
}
