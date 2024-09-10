package router

import (
	"github.com/gin-gonic/gin"
	"visual-state-machine/config"
	"visual-state-machine/internal/api"
)

type Router struct {
	engine   *gin.Engine
	register *Register
}

func NewRouter() *Router {
	return &Router{
		engine:   gin.Default(),
		register: newRegister(),
	}
}

func (rt *Router) InitRouter() error {
	gin.SetMode(config.Get().GinMode)
	router := rt.engine
	apis := api.InitApis()
	rt.register.r = router
	rt.register.apis = apis
	rt.register.registerRouter()
	// 启动服务
	err := router.Run(":" + config.Get().Port)
	if err != nil {
		return err
	}
	return nil
}
