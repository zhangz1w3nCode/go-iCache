package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/MoeGolibrary/go-lib/zlog"

	"visual-state-machine/config"
	user_svc "visual-state-machine/internal/api/user"
	user_logic "visual-state-machine/internal/logic/user"
	user_repo "visual-state-machine/internal/repo/user"
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

	userRepo := user_repo.New()
	userLogic := user_logic.New(userRepo)
	userService := user_svc.NewService(userLogic)

	// 设置 HTTP 路由
	http.HandleFunc("/user/get", userService.Get)
	// 启动 HTTP 服务器
	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("server stopped")
}
