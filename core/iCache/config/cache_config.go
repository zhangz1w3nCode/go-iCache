package config

import (
	"time"
)

// GoCacheConfig 缓存配置
type GoCacheConfig struct {
	//缓存名称
	CacheName string
	//缓存类型
	CacheType string
	//过期时间
	ExpireTime time.Duration
	//清理时间
	CleanTime time.Duration
}

// RistrettoCacheConfig 缓存配置
type RistrettoCacheConfig struct {
	//缓存名称
	CacheName string
	//缓存类型
	CacheType   string
	NumCounters int64
	MaxCost     int64
	BufferItems int64
	Metrics     bool
}
