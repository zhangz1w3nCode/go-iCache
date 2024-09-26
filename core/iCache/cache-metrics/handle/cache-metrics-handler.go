package Handler

import (
	cacheMcs "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-metrics/metrics"
	"sync"
	"time"
)

// MetricHandler 指标收集器
type MetricHandler struct {
	mu          sync.Mutex
	history     []*cacheMcs.CacheMetrics
	capacity    int
	currentSize int64
	cycle       time.Duration
}

// NewMetricHandler 新建指标收集器
func NewMetricHandler(cycle time.Duration, capacity int) *MetricHandler {
	return &MetricHandler{
		history:     make([]*cacheMcs.CacheMetrics, 0, capacity),
		capacity:    capacity,
		cycle:       cycle,
		currentSize: 0,
	}
}

// Collect 收集指标
func (c *MetricHandler) Collect(metric *cacheMcs.CacheMetrics) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.history = append(c.history, metric)
	if len(c.history) > c.capacity {
		c.history = c.history[1:]
	}
	c.currentSize = metric.CacheSize
}
