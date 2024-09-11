package router

import (
	"github.com/gin-gonic/gin"
	"visual-state-machine/internal/api"
)

type Register struct {
	engine *gin.Engine
	apis   *api.Apis
}

func newRegister(engine *gin.Engine, apis *api.Apis) *Register {
	return &Register{
		engine: engine,
		apis:   apis,
	}
}

func (rgst *Register) registerRouter() {
	rgst.registerUserRouter()
	rgst.registerFlowRouter()
	//additional register
}
