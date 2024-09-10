package router

import (
	"github.com/gin-gonic/gin"
	"visual-state-machine/config"
	"visual-state-machine/internal/api"
)

type Router struct {
	engine *gin.Engine
}

func InitRouter() error {
	gin.SetMode(config.Get().GinMode)
	router := gin.Default()
	apis := api.InitApis()
	// 注册路由
	registerRouter(router, apis)
	// 启动服务
	err := router.Run(":" + config.Get().Port)
	if err != nil {
		return err
	}
	return nil
}
