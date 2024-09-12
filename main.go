package main

import (
	"fmt"
	core "github.com/zhangz1w3nCode/go-iCache/core/iCache"
	"time"
)

func main() {
	caffeineCache := core.NewCaffeineCache("test_cache", &core.CacheConfig{
		ExpireAfterWrite:  5 * time.Minute,
		ExpireAfterAccess: 10 * time.Minute,
	})

	get := caffeineCache.Get("key1")

	if get != nil {
		fmt.Println(get)
	} else {
		fmt.Println("not found reload")
		caffeineCache.Set("key1", "value1")
	}

	value1 := caffeineCache.Get("key1")
	fmt.Println(value1)
}
