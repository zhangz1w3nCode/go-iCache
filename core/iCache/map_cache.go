package iCache

import (
	"sync"
)

// SimpleCache 简单的缓存实现
type SimpleCache struct {
	cache map[string]*ValueWrapper
	lock  sync.RWMutex
	name  string
}

// NewSimpleCache 创建一个新的SimpleCache实例
func NewSimpleCache(name string) *SimpleCache {
	return &SimpleCache{
		cache: make(map[string]*ValueWrapper),
		name:  name,
	}
}

func (c *SimpleCache) Get(key string) *ValueWrapper {
	c.lock.RLock()
	defer c.lock.RUnlock()
	vw, exists := c.cache[key]
	if exists {
		vw.UpdateAccessTime()
	}
	return vw
}

func (c *SimpleCache) Put(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache[key] = NewValueWrapper(value)
}

func (c *SimpleCache) GetValues() []*ValueWrapper {
	c.lock.RLock()
	defer c.lock.RUnlock()
	var values []*ValueWrapper
	for _, vw := range c.cache {
		values = append(values, vw)
	}
	return values
}

func (c *SimpleCache) GetKeys() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	var keys []string
	for key := range c.cache {
		keys = append(keys, key)
	}
	return keys
}

func (c *SimpleCache) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.cache)
}

func (c *SimpleCache) GetName() string {
	return c.name
}

func (c *SimpleCache) CalculateMemoryUsage() float64 {
	// 简化版本，不计算实际内存使用
	return float64(c.Size())
}

func (c *SimpleCache) GetCacheStatus() CacheStats {
	// 简化版本，不提供真实缓存状态
	return CacheStats{}
}
