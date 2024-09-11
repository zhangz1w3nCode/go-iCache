package router

func (rgst *Register) registerFlowRouter() {
	//flowAPI
	userAPI := rgst.apis.UserApi

	//URL Mapping API
	userGroup := rgst.engine.Group("/flow")
	{
		userGroup.GET("/get/:id", userAPI.Get)
	}
}
