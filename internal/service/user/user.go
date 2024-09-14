package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/user"
	"github.com/zhangz1w3nCode/go-iCache/internal/logic/cache"
)

type UserService struct {
	user.UnimplementedUserServiceServer
	logic   *cache.TestLogic
	manager *manager.CacheManager
}

func NewUserService(mm *manager.CacheManager) *UserService {
	return &UserService{
		logic:   cache.NewTestLogic(),
		manager: mm,
	}
}

func (s *UserService) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {

	icache := s.manager.GetCache("user_cache")

	value := icache.Get(in.GetUserID())

	return &user.GetUserResponse{UserName: value.Data.(string)}, nil
}
