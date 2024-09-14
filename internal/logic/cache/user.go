package cache

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/config"
	go_cache "github.com/zhangz1w3nCode/go-iCache/core/iCache/go-cache"
	"log"
	"time"
)

type TestLogic struct {
	cache *go_cache.GoCache
}

func NewTestLogic() *TestLogic {
	return &TestLogic{
		cache: go_cache.NewGoCache(&config.GoCacheConfig{
			CacheName:  "TestCache",
			CacheType:  "goCache",
			ExpireTime: 5 * time.Minute,
			CleanTime:  10 * time.Minute,
		}),
	}
}

func (u *TestLogic) GetCache(ctx context.Context, key string) string {
	log.Println("GetCache:key" + key)
	return u.cache.Get("user_id_" + key).Data.(string)
}

func (u *TestLogic) SetCache(ctx context.Context, key string, value string) {
	log.Println("SetCache:key:" + key + "value:" + value)
	u.cache.Set("user_id_"+key, value)
}
