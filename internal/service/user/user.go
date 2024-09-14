package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/user"
	"github.com/zhangz1w3nCode/go-iCache/internal/logic/cache"
)

type UserService struct {
	user.UnimplementedUserServiceServer
	logic *cache.TestLogic
}

func NewUserService() *UserService {
	return &UserService{
		logic: cache.NewTestLogic(),
	}
}

func (s *UserService) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {

	userCache := manager.NewCacheManager().GetCache("user_cache")

	value := userCache.Get(in.GetUserID())

	return &user.GetUserResponse{UserName: value.Data.(string)}, nil
}
