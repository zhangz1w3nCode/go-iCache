package cacheRegister

import (
	"github.com/zhangz1w3nCode/go-iCache/core/etcd"
	cacheInit "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-init"
	monitorsvc "github.com/zhangz1w3nCode/go-iCache/internal/service/monitor"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc"
	"log"
)

func RegisterCacheGrpcService(s *grpc.Server, serviceName, bizAppIp string, etcdAddress string, cacheInit *cacheInit.CacheInit) *etcd.EtcdRegister {
	monitorService := monitorsvc.NewMonitorService(cacheInit.CacheManager)
	monitorpb.RegisterCacheMonitorServiceServer(s, monitorService)
	etcdRegister, err := etcd.NewEtcdRegister(etcdAddress)
	if err != nil {
		log.Fatalf("failed to get etcd connect: %v", err)
	}
	err = etcdRegister.ServiceRegister("/services/"+serviceName+"/"+bizAppIp, bizAppIp, 60)
	if err != nil {
		log.Fatalf("failed to cache-register etcd: %v", err)
	}
	return etcdRegister
}
func CleanUpCacheGrpcService(etcd *etcd.EtcdRegister) {
	err := etcd.Close()
	if err != nil {
		log.Printf("failed to clean up etcd: %v", err)
	}
}
