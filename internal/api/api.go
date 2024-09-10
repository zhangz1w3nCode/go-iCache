package api

import (
	userApi "visual-state-machine/internal/api/user"
)

// Apis 是所有Api的集合
type Apis struct {
	UserApiService *userApi.Api
}

func newApis(
	userApiService *userApi.Api,
) *Apis {
	return &Apis{
		UserApiService: userApiService,
	}
}
