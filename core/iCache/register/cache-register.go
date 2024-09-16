package register

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/zhangz1w3nCode/go-iCache/config"
	start "github.com/zhangz1w3nCode/go-iCache/core/iCache/start"
	monitorsvc "github.com/zhangz1w3nCode/go-iCache/internal/service/monitor"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"time"
)

func RegisterZookeeper(zookeeperServers []string, serviceName string, ip string, info map[string]grpc.ServiceInfo) error {
	zkConn, _, err := zk.Connect(zookeeperServers, time.Second*10)
	if err != nil {
		return err
	}
	defer func() {
		zkConn.Close()
	}()

	path0 := "/services"
	if _, err := zkConn.Create(path0, nil, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNoNode {
			return err
		}
	}

	path := "/services/" + serviceName
	if _, err := zkConn.Create(path, nil, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNoNode {
			return err
		}
	}
	path2 := "/services/" + serviceName + "/" + ip
	infoStr, err := json.Marshal(info)
	data := infoStr
	if _, err := zkConn.Create(path2, data, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNoNode {
			return err
		}
	}

	return nil
}

func RegisterCacheServcie(s *grpc.Server, serviceName, bizAppIp string, zkIp string) {
	//读取配置文件地址设置
	config.Init("./config/config.yaml")

	managerCache := start.NewCacheInit().CacheManager
	monitorService := monitorsvc.NewMonitorService(managerCache)
	monitorpb.RegisterCacheMonitorServiceServer(s, monitorService)
	grpc_health_v1.RegisterHealthServer(s, monitorService)

	info := s.GetServiceInfo()

	errZk := RegisterZookeeper([]string{zkIp}, serviceName, bizAppIp, info)

	if errZk != nil {
		log.Fatalf("failed to start zookeeper: %v", errZk)
	}

	go func() {
		<-context.Background().Done()
		fmt.Println("关闭一些资源 比如zk")
		s.GracefulStop()
	}()
}
