package valueWrapper

import (
	cacheMetrics "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/cache-metrics"
	"sync/atomic"
	"time"
	"unsafe"
)

// ValueWrapper 缓存的数据包装类
type ValueWrapper struct {
	Data              interface{}                     // 缓存的数据
	WriteTime         int64                           // 缓存写入时间
	AccessTime        int64                           // 缓存最后访问时间
	CacheValueMetrics *cacheMetrics.CacheValueMetrics // 缓存状态
}

// NewValueWrapper 创建一个新的ValueWrapper实例
func NewValueWrapper(data interface{}) *ValueWrapper {
	return &ValueWrapper{
		Data:              data,
		WriteTime:         time.Now().Unix(),
		AccessTime:        time.Now().Unix(),
		CacheValueMetrics: cacheMetrics.NewCacheValueMetrics(),
	}
}

// UpdateAccessTime 更新最后访问时间
func (vw *ValueWrapper) UpdateAccessTime() {
	vw.AccessTime = time.Now().Unix()
}

// UpdateWriteTime 更新最后写入时间
func (vw *ValueWrapper) UpdateWriteTime() {
	vw.WriteTime = time.Now().Unix()
}

// UpdateCacheValueMetrics 更新缓存状态
func (vw *ValueWrapper) UpdateCacheValueMetrics() {
	currentStatus := vw.CacheValueMetrics
	atomic.AddInt64(&currentStatus.CacheHitCount, 1)
	currentStatus.CacheSize = float32(unsafe.Sizeof(vw.Data))
}
