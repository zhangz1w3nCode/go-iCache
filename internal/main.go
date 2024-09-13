package main

import (
	"flag"
	"github.com/zhangz1w3nCode/go-iCache/config"
	pb "github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	svc "github.com/zhangz1w3nCode/go-iCache/internal/service/test"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	configPath := flag.String("config", "", "specify config path [config.yaml]")
	flag.Parse()
	if configPath == nil || *configPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	config.Init(*configPath)

	listen, _ := net.Listen("tcp", ":"+config.Get().Port)

	s := grpc.NewServer()

	testService := svc.NewTestService()
	pb.RegisterTestServiceServer(s, testService)

	log.Printf("server start in port: %s", config.Get().Port)

	err := s.Serve(listen)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
