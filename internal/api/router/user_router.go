package router

import (
	"github.com/gin-gonic/gin"
	"visual-state-machine/internal/api"
)

type Register struct {
	r    *gin.Engine
	apis *api.Apis
}

func newRegister() *Register {
	return &Register{}
}

func (rgst *Register) registerRouter() {
	rgst.registerUserRouter()
	rgst.registerFlowRouter()
}

func (rgst *Register) registerUserRouter() {
	//userAPI
	userAPI := rgst.apis.UserApi

	//URL Mapping API
	userGroup := rgst.r.Group("/user")
	{
		userGroup.GET("/get/:id", userAPI.Get)
	}
}

func (rgst *Register) registerFlowRouter() {
	//userAPI
	userAPI := rgst.apis.UserApi

	//URL Mapping API
	userGroup := rgst.r.Group("/flow")
	{
		userGroup.GET("/get/:id", userAPI.Get)
	}
}
