package register

import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
	SS "github.com/zhangz1w3nCode/go-iCache/core/iCache/start"
	userpb "github.com/zhangz1w3nCode/go-iCache/internal/api/generate/user"
	usersvc "github.com/zhangz1w3nCode/go-iCache/internal/service/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	CacheAPI *SS.CacheApi
)

func RegisterZookeeper(zookeeperServers []string, serviceName string, ip string, info map[string]grpc.ServiceInfo) error {
	zkConn, _, err := zk.Connect(zookeeperServers, time.Second*10)
	if err != nil {
		return err
	}
	defer zkConn.Close()

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
	managerCache := SS.NewCacheInit().CacheManager
	userService := usersvc.NewUserService(managerCache)
	userpb.RegisterUserServiceServer(s, userService)

	CacheAPI = SS.NewCacheApi(managerCache)

	info := s.GetServiceInfo()

	errZk := RegisterZookeeper([]string{zkIp}, serviceName, bizAppIp, info)

	if errZk != nil {
		log.Fatalf("failed to start zookeeper: %v", errZk)
	}
}
