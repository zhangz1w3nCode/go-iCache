package router

func (rgst *Register) registerUserRouter() {
	//userAPI
	userAPI := rgst.apis.UserApi

	//URL Mapping API
	userGroup := rgst.engine.Group("/user")
	{
		userGroup.GET("/get/:id", userAPI.Get)
		userGroup.POST("/list", userAPI.List)
		//userGroup.POST("/create", userAPI.Create)
		//userGroup.POST("/update", userAPI.Update)
		//userGroup.POST("/delete", userAPI.Delete)
	}
}
