package cacheRegister

import (
	"github.com/zhangz1w3nCode/go-iCache/core/etcd"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	monitorsvc "github.com/zhangz1w3nCode/go-iCache/internal/service/monitor"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc"
	"log"
)

func RegisterCacheGrpcService(s *grpc.Server, serviceName, bizAppIp string, etcdAddress string, managerCache *manager.CacheManager) *etcd.EtcdRegister {
	monitorService := monitorsvc.NewMonitorService(managerCache)
	monitorpb.RegisterCacheMonitorServiceServer(s, monitorService)
	etcdRegister, err := etcd.NewEtcdRegister(etcdAddress)
	if err != nil {
		log.Fatalf("failed to get etcd connect: %v", err)
	}
	err = etcdRegister.ServiceRegister("/services/"+serviceName+"/"+bizAppIp, bizAppIp, 60)
	if err != nil {
		log.Fatalf("failed to register etcd: %v", err)
	}
	return etcdRegister
}
func CleanUpCacheGrpcService(etcd *etcd.EtcdRegister) {
	err := etcd.Close()
	if err != nil {
		log.Printf("failed to clean up etcd: %v", err)
	}
}
