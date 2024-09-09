package main

import (
	"flag"
	"fmt"
	"github.com/MoeGolibrary/go-lib/zlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"visual-state-machine/config"
	"visual-state-machine/internal/service/flow"
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

	log.Printf("server started at :%d, debug: %t", config.Get().Port, config.Get().Debug)

	// 设置 HTTP 路由
	http.HandleFunc("/generate_fsm", flow.GenerateFsmHandler)
	// 启动 HTTP 服务器
	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("server stopped")
}
