package api

import (
	userApi "visual-state-machine/internal/api/user"
)

// Apis 是所有Api的集合
type Apis struct {
	UserApi *userApi.Api
}

func newApis(
	userApi *userApi.Api,
) *Apis {
	return &Apis{
		UserApi: userApi,
	}
}
