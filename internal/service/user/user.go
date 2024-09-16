package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	user2 "github.com/zhangz1w3nCode/go-iCache/pb/generate/user"
)

type UserService struct {
	user2.UnimplementedUserServiceServer
	manager *manager.CacheManager
}

func NewUserService(mm *manager.CacheManager) *UserService {
	return &UserService{
		manager: mm,
	}
}

func (s *UserService) GetUser(ctx context.Context, in *user2.GetUserRequest) (*user2.GetUserResponse, error) {

	icache := s.manager.GetCache("user_cache")

	value := icache.Get(in.GetUserID())

	return &user2.GetUserResponse{UserName: value.Data.(string)}, nil
}
