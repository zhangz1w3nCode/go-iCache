package router

import (
	"github.com/gin-gonic/gin"
	"visual-state-machine/internal/api"
)

func registerRouter(r *gin.Engine, apis *api.Apis) {
	registerUserRouter(r, apis)
	registerFlowRouter(r, apis)
}

func registerUserRouter(r *gin.Engine, apis *api.Apis) {
	//userAPI
	userAPI := apis.UserApi

	//URL Mapping API
	userGroup := r.Group("/user")
	{
		userGroup.GET("/get/:id", userAPI.Get)
	}
}

func registerFlowRouter(r *gin.Engine, apis *api.Apis) {
	////FlowAPI
	//userAPI := apis.UserApi
	//
	////URL Mapping API
	//userGroup := r.Group("/user")
	//{
	//	userGroup.GET("/get/:id", userAPI.Get)
	//}
}
