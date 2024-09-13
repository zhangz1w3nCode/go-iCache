package value_wrapper

import (
	"time"
)

// ValueWrapper 缓存的数据包装类
type ValueWrapper struct {
	Data       interface{} // 缓存的数据
	WriteTime  int64       // 缓存写入时间
	AccessTime int64       // 缓存最后访问时间
}

// NewValueWrapper 创建一个新的ValueWrapper实例
func NewValueWrapper(data interface{}) *ValueWrapper {
	return &ValueWrapper{
		Data:       data,
		WriteTime:  time.Now().Unix(),
		AccessTime: time.Now().Unix(),
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
