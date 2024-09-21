package monitor

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	monitorlogic "github.com/zhangz1w3nCode/go-iCache/internal/logic/monitor"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
)

type MonitorService struct {
	monitorpb.UnimplementedCacheMonitorServiceServer
	monitorLogic *monitorlogic.MonitorLogic
}

func NewMonitorService(manager *manager.CacheManager) *MonitorService {
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
