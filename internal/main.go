package main

import (
	"flag"
	"github.com/MoeGolibrary/go-lib/zlog"
	"log"
	"os"
	"os/signal"
	"syscall"
	"visual-state-machine/config"
	"visual-state-machine/internal/router"
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

	//初始化路由
	router.InitRouter()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("server stopped")
}
