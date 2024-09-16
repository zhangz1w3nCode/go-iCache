package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/user"
)

type UserService struct {
	user.UnimplementedUserServiceServer
	manager *manager.CacheManager
}

func NewUserService(mm *manager.CacheManager) *UserService {
	return &UserService{
		manager: mm,
	}
}

func (s *UserService) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {

	icache := s.manager.GetCache("user_cache")

	value := icache.Get(in.GetUserID())

	return &user.GetUserResponse{UserName: value.Data.(string)}, nil
}
