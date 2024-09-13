package iCache

import (
	"time"
)

// CacheConfig 缓存配置
type CacheConfig struct {
	//缓存名称
	CacheName string
	//缓存类型
	CacheType string
	//过期时间
	ExpireTime time.Duration
	//清理时间
	CleanTime time.Duration
}
