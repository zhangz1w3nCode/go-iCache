package router

import (
	"context"
	"github.com/MoeGolibrary/go-lib/zlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"visual-state-machine/config"
	"visual-state-machine/internal/api"
)

type Router struct {
	engine   *gin.Engine
	register *Register
}

func NewRouter(apis *api.Apis) *Router {
	engine := gin.Default()
	return &Router{
		engine:   engine,
		register: newRegister(engine, apis),
	}
}

func (rt *Router) Init() error {
	// 设置模式：是否是debug模式
	gin.SetMode(config.Get().GinMode)
	// 注册路由
	rt.register.registerRouter()
	// 启动服务
	err := rt.engine.Run(":" + config.Get().Port)
	if err != nil {
		return err
	}
	return nil
}

func InitRouter() {
	apis := api.InitApis()
	router := NewRouter(apis)
	err := router.Init()
	if err != nil {
		zlog.Error(context.Background(), "InitRouter Error!", zap.Error(err))
	}
}
