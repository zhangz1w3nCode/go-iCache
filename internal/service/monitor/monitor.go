package monitor

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	monitorlogic "github.com/zhangz1w3nCode/go-iCache/internal/logic/monitor"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type MonitorService struct {
	monitorpb.UnimplementedCacheMonitorServiceServer
	grpc_health_v1.UnimplementedHealthServer
	monitorLogic *monitorlogic.MonitorLogic
}

func NewMonitorService(manager *manager.CacheManager) *MonitorService {
	return &MonitorService{
		monitorLogic: monitorlogic.NewMonitorService(manager),
	}
}

func (m *MonitorService) GetCacheUserIpList(ctx context.Context, req *monitorpb.GetCacheUserIpListRequest) (*monitorpb.GetCacheUserIpListResponse, error) {

	ipList, err := m.monitorLogic.GetCacheUserIpList(ctx, req.GetAppName())

	if err != nil {
		return &monitorpb.GetCacheUserIpListResponse{UserCacheIpList: []string{}}, nil
	}

	return &monitorpb.GetCacheUserIpListResponse{UserCacheIpList: ipList}, nil
}
func (m *MonitorService) GetCacheUserAppNameList(ctx context.Context, req *monitorpb.GetCacheUserAppNameListRequest) (*monitorpb.GetCacheUserAppNameListResponse, error) {

	AppNameList, err := m.monitorLogic.GetCacheUserAppNameList(ctx)

	if err != nil {
		return &monitorpb.GetCacheUserAppNameListResponse{UserCacheAppNameList: []string{}}, nil
	}

	return &monitorpb.GetCacheUserAppNameListResponse{UserCacheAppNameList: AppNameList}, nil
}
func (m *MonitorService) GetCacheNameList(ctx context.Context, req *monitorpb.GetCacheNameListRequest) (*monitorpb.GetCacheNameListResponse, error) {
	return &monitorpb.GetCacheNameListResponse{CacheNameList: make([]string, 0)}, nil
}

func (m *MonitorService) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (m *MonitorService) Watch(in *grpc_health_v1.HealthCheckRequest, srv grpc_health_v1.Health_WatchServer) error {
	return nil
}
