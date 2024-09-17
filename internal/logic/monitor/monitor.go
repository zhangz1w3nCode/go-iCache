package test

import (
	"context"
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

// GetCacheUserAddressList ：获取使用了iCache的机器的ip集合:哪个appName对应的哪些机器用了iCache
func (m *MonitorLogic) GetCacheUserAddressList(ctx context.Context, appName string) ([]string, error) {
	zkAddress := config.Get().ZkAddress

	if zkAddress == "" {
		return nil, status.Errorf(codes.Unavailable, "Get zookeeper ip error!")
	}

	zkConn, _, err := zk.Connect([]string{zkAddress}, time.Second*10)
	if err != nil {
		return nil, err
	}

	path := "/services"

	exists, stat, err := zkConn.Exists(path)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path+" from zookeeper error!")
	}
	if stat == nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path+" from zookeeper stat error!")
	}
	if !exists {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path+" from zookeeper not exists!")
	}

	path1 := path + "/" + appName

	exists, stat, err = zkConn.Exists(path1)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path1+" from zookeeper error!")
	}
	if stat == nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path1+" from zookeeper stat error!")
	}
	if !exists {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path1+" from zookeeper not exists!")
	}

	addressList, stat, err := zkConn.Children(path1)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path1+" from zookeeper error!")
	}
	if stat == nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource "+path1+" from zookeeper stat error!")
	}

	return addressList, nil
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

	exists, stat, err := zkConn.Exists("/services")

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource from zookeeper error!")
	}
	if stat == nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource from zookeeper stat error!")
	}
	if !exists {
		return nil, status.Errorf(codes.Unavailable, "Get path resource from zookeeper not exists!")
	}

	appNameList, stat1, err1 := zkConn.Children("/services")

	if err1 != nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource from zookeeper error!")
	}
	if stat1 == nil {
		return nil, status.Errorf(codes.Unavailable, "Get path resource from zookeeper stat error!")
	}

	return appNameList, nil
}

// GetCacheNameList 获取某个appName下的某台机器上有哪些缓存：如productCache、userCache
func (m *MonitorLogic) GetCacheNameList() ([]string, error) {
	cacheName := m.manager.GetAllCacheName()
	return cacheName, nil
}
