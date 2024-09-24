package value_wrapper

import (
	"time"
)

// ValueWrapper 缓存的数据包装类
type ValueWrapper struct {
	Data        interface{} // 缓存的数据
	WriteTime   int64       // 缓存写入时间
	AccessTime  int64       // 缓存最后访问时间
	CacheStatus *CacheStatus
}

type CacheStatus struct {
	CacheCurrentSize int64
	CacheMaxSize     int64
	CacheHit         int64
	CacheMiss        int64
	CacheQuery       int64
}

// NewValueWrapper 创建一个新的ValueWrapper实例
func NewValueWrapper(data interface{}) *ValueWrapper {
	return &ValueWrapper{
		Data:       data,
		WriteTime:  time.Now().Unix(),
		AccessTime: time.Now().Unix(),
		CacheStatus: &CacheStatus{
			CacheCurrentSize: 0,
			CacheMaxSize:     0,
			CacheHit:         0,
			CacheMiss:        0,
			CacheQuery:       0,
		},
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

// UpdateCacheStatus 更新缓存状态
func (vw *ValueWrapper) UpdateCacheStatus(cacheStatus *CacheStatus) {
	vw.CacheStatus = cacheStatus
}
