package test

import (
	"context"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/zhangz1w3nCode/go-iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
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

	zkAddress := config.Get().ZkAddress

	if zkAddress == "" {
		return nil, status.Errorf(codes.Unavailable, "Get zookeeper ip error!")
	}

	zkConn, _, err := zk.Connect([]string{zkAddress}, time.Second*10)
	if err != nil {
		return nil, err
	}
	defer func() {
		fmt.Println("关闭zk一些资源")
		zkConn.Close()
	}()

	path := "/services"

	resource, stat, err := zkConn.Get(path)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource from zookeeper error!")
	}
	if stat == nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource from zookeeper stat error!")
	}

	fmt.Println(resource)
	fmt.Println(stat)

	return nil, status.Errorf(codes.Unimplemented, "method GetCacheUserAppNameList not implemented")
}
func (m *MonitorLogic) GetCacheNameList(ctx context.Context, appName string) ([]string, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheNameList not implemented")
}
