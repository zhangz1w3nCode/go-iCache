package main

import (
	"flag"
	"fmt"
	"github.com/MoeGolibrary/go-lib/server"
	"github.com/MoeGolibrary/go-lib/zlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"visual-state-machine/config"
	"visual-state-machine/internal/utils/myFsm"
)

func main() {
	// 这一堆都应该放到go-lib里，拆分为 NewServer, RegisterService, StartAndStop
	// 初始化配置
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	flag.Parse()
	if configPath == nil || *configPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	config.Init(*configPath)

	// 初始化日志
	zlog.InitLogger(zlog.NewConfig())

	// 初始化 server 和 services
	s := server.NewDefaultServer()
	//services := service.InitServices()

	// 使用特定的服务注册函数来检查服务是否已实现
	//engagementsvcpb.RegisterSettingServiceServer(s, services.SettingService)

	s.Start()
	log.Printf("server started at :%d, debug: %t", config.Get().Port, config.Get().Debug)

	// 设置 HTTP 路由
	http.HandleFunc("/generate_fsm", myFsm.GenerateFsmHandler)

	// 启动 HTTP 服务器
	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	s.Stop()
	log.Println("server stopped")
}
