package register

import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
	testpb "github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	userpb "github.com/zhangz1w3nCode/go-iCache/internal/api/generate/user"
	testsvc "github.com/zhangz1w3nCode/go-iCache/internal/service/test"
	usersvc "github.com/zhangz1w3nCode/go-iCache/internal/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

//func main() {
//	StartGRPCServer("iCacheService", "192.168.31.84:9090", "192.168.31.84:2181")
//}

func RegisterService(zookeeperServers []string, serviceName string, ip string, info map[string]grpc.ServiceInfo) error {
	zkConn, _, err := zk.Connect(zookeeperServers, time.Second*10)
	if err != nil {
		return err
	}
	defer zkConn.Close()

	path := "/services/" + serviceName + "/" + ip
	infoStr, err := json.Marshal(info)
	data := infoStr
	if _, err := zkConn.Create(path, data, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err != zk.ErrNodeExists {
			return err
		}
	}

	return nil
}

func StartGRPCServer(serviceName, bizAppIp string, zkIp string) {
	listen, _ := net.Listen("tcp", bizAppIp)
	s := grpc.NewServer()
	reflection.Register(s)
	testService := testsvc.NewTestService()
	userService := usersvc.NewUserService()
	testpb.RegisterTestServiceServer(s, testService)
	userpb.RegisterUserServiceServer(s, userService)
	info := s.GetServiceInfo()

	errZk := RegisterService([]string{zkIp}, serviceName, bizAppIp, info)

	if errZk != nil {
		log.Fatalf("failed to start zookeeper: %v", errZk)
		return
	}

	err := s.Serve(listen)

	log.Printf("server start in port: %s", bizAppIp)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
