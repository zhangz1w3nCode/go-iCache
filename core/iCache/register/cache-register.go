package register

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	pb "github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	svc "github.com/zhangz1w3nCode/go-iCache/internal/service/test"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

func RegisterService(zookeeperServers []string, serviceName, serviceAddress string) error {
	zkConn, _, err := zk.Connect(zookeeperServers, time.Second*10)
	if err != nil {
		return err
	}
	defer zkConn.Close()

	path := "/" + serviceName
	ips := strings.Join(GetIPs(), ",")
	data := []byte("/" + serviceName + "/" + ips)
	if _, err := zkConn.Create(path, data, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err != zk.ErrNodeExists {
			return err
		}
	}

	// 读取节点数据
	get, _, err := zkConn.Get(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Node data:", string(get))

	return nil
}

func StartGRPCServer(serviceName, serviceAddress string) {

	if err := RegisterService([]string{"192.168.31.84:2181"}, serviceName, serviceAddress); err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	listen, _ := net.Listen("tcp", ":9099")
	s := grpc.NewServer()
	testService := svc.NewTestService()
	pb.RegisterTestServiceServer(s, testService)

	log.Printf("server start in port: %s", "9099")

	err := s.Serve(listen)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}

func GetIPs() []string {
	// 获取本机的IP地址
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Error getting interface addresses:", err)
		return nil
	}

	var ipList []string

	for _, addr := range addrs {
		// 检查地址是否是IPv4或IPv6
		ip, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		// 检查是否是回环地址（127.0.0.1）
		if ip.IP.IsLoopback() {
			continue
		}

		// 打印IPv4地址
		if ip.IP.To4() != nil {
			ipList = append(ipList, ip.IP.String())
		}
	}

	return ipList
}
