package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"visual-state-machine/internal/api"

	"github.com/MoeGolibrary/go-lib/zlog"

	"visual-state-machine/config"
)

func main() {
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

	//初始化所有api
	Apis := api.InitApis()

	// 设置 HTTP 路由
	http.HandleFunc("/user/get", Apis.UserApiService.Get)

	//绑定端口
	http.ListenAndServe(":"+config.Get().Port, nil)
	log.Printf("server started at :%s, debug: %t", config.Get().Port, config.Get().Debug)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("server stopped")
}
