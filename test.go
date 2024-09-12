package main

import (
	"fmt"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache"
)

func main() {
	// 创建缓存实例
	cacheInstance := iCache.NewSimpleCache("testCache")

	for {
		// 获取数据
		valueWrapper := cacheInstance.Get("key1")
		if valueWrapper != nil {
			fmt.Println("cache value:", valueWrapper.Data)
		} else {
			// 存储数据
			cacheInstance.Put("key1", "db")
		}

		// 获取所有键
		keys := cacheInstance.GetKeys()
		fmt.Println("Keys:", keys)

		// 获取所有值
		values := cacheInstance.GetValues()
		fmt.Println("Values:", values)

		// 获取缓存大小
		size := cacheInstance.Size()
		fmt.Println("Cache size:", size)

		// 获取缓存名称
		name := cacheInstance.GetName()
		fmt.Println("Cache name:", name)

		// 计算内存使用情况
		memoryUsage := cacheInstance.CalculateMemoryUsage()
		fmt.Println("Memory usage:", memoryUsage)

		// 获取缓存状态
		cacheStatus := cacheInstance.GetCacheStatus()
		fmt.Println("Cache status:", cacheStatus)
	}
}
