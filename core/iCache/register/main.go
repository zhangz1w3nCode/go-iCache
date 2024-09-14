package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	pb "github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	svc "github.com/zhangz1w3nCode/go-iCache/internal/service/test"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func RegisterService(zookeeperServers []string, serviceName, serviceAddress string) error {
	zkConn, _, err := zk.Connect(zookeeperServers, time.Second*10)
	if err != nil {
		return err
	}
	defer zkConn.Close()

	path := "/services/" + serviceName
	if _, err := zkConn.Create(path, []byte(path), int32(0), zk.WorldACL(zk.PermAll)); err != nil {
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
