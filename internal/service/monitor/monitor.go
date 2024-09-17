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

func (m *MonitorService) GetCacheUserAddressList(ctx context.Context, req *monitorpb.GetCacheUserAddressListRequest) (*monitorpb.GetCacheUserAddressListResponse, error) {

	addressList, err := m.monitorLogic.GetCacheUserAddressList(ctx, req.GetAppName())

	if err != nil {
		return &monitorpb.GetCacheUserAddressListResponse{UserCacheAddressList: []string{}}, nil
	}

	return &monitorpb.GetCacheUserAddressListResponse{UserCacheAddressList: addressList}, nil
}
func (m *MonitorService) GetCacheUserAppNameList(ctx context.Context, req *monitorpb.GetCacheUserAppNameListRequest) (*monitorpb.GetCacheUserAppNameListResponse, error) {

	AppNameList, err := m.monitorLogic.GetCacheUserAppNameList(ctx)

	if err != nil {
		return &monitorpb.GetCacheUserAppNameListResponse{UserCacheAppNameList: []string{}}, nil
	}

	return &monitorpb.GetCacheUserAppNameListResponse{UserCacheAppNameList: AppNameList}, nil
}
func (m *MonitorService) GetCacheNameList(ctx context.Context, req *monitorpb.GetCacheNameListRequest) (*monitorpb.GetCacheNameListResponse, error) {
	cacheNameList, err := m.monitorLogic.GetCacheNameList()

	if err != nil {
		return &monitorpb.GetCacheNameListResponse{CacheNameList: []string{}}, nil
	}

	return &monitorpb.GetCacheNameListResponse{CacheNameList: cacheNameList}, nil
}
