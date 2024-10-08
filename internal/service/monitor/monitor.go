package monitor

import (
	"context"
	cacheManager "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	monitorlogic "github.com/zhangz1w3nCode/go-iCache/internal/logic/monitor"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
)

type MonitorService struct {
	monitorpb.UnimplementedCacheMonitorServiceServer
	monitorLogic *monitorlogic.MonitorLogic
}

func NewMonitorService(manager *cacheManager.CacheManager) *MonitorService {
	return &MonitorService{
		monitorLogic: monitorlogic.NewMonitorService(manager),
	}
}

func (m *MonitorService) GetCacheNameList(ctx context.Context, req *monitorpb.GetCacheNameListRequest) (*monitorpb.GetCacheNameListResponse, error) {
	cacheNameList, err := m.monitorLogic.GetCacheNameList()

	if err != nil {
		return nil, err
	}

	return &monitorpb.GetCacheNameListResponse{CacheNameList: cacheNameList}, nil
}

func (m *MonitorService) GetCacheKeyList(ctx context.Context, req *monitorpb.GetCacheKeyListRequest) (*monitorpb.GetCacheKeyListResponse, error) {

	cacheNameList, err := m.monitorLogic.GetCacheKeyList(req.GetCacheName())

	if err != nil {
		return nil, err
	}

	return &monitorpb.GetCacheKeyListResponse{CacheKeyList: cacheNameList}, nil
}
func (m *MonitorService) GetValueToCacheUser(ctx context.Context, req *monitorpb.GetValueToCacheUserRequest) (*monitorpb.GetValueToCacheUserResponse, error) {
	cacheNameList, err := m.monitorLogic.GetValueToCacheUser(req.GetCacheName(), req.GetCacheKey())

	if err != nil {
		return nil, err
	}

	return &monitorpb.GetValueToCacheUserResponse{CacheValue: cacheNameList}, nil
}

func (m *MonitorService) GetCacheMetrics(ctx context.Context, req *monitorpb.GetCacheMetricsRequest) (*monitorpb.GetCacheMetricsResponse, error) {
	cacheMetrics, err := m.monitorLogic.GetCacheMetrics(req.GetCacheName())

	if err != nil {
		return nil, err
	}

	return &monitorpb.GetCacheMetricsResponse{CacheMetrics: cacheMetrics}, nil
}
