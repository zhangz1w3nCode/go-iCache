package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MonitorLogic struct {
	monitor monitorpb.UnimplementedCacheMonitorServiceServer
	manager *manager.CacheManager
}

func NewMonitorService(manager *manager.CacheManager) *MonitorLogic {
	return &MonitorLogic{
		manager: manager,
	}
}

func (m *MonitorLogic) GetCacheUserIpList(ctx context.Context, appName string) ([]string, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheUserIpList not implemented")
}
func (m *MonitorLogic) GetCacheUserAppNameList(ctx context.Context) ([]string, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheUserAppNameList not implemented")
}
func (m *MonitorLogic) GetCacheNameList(ctx context.Context, appName string) ([]string, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheNameList not implemented")
}
