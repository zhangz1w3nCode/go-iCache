package test

import (
	"context"
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
	s.logic.SetCache(ctx, in.GetUserID(), "testSet")
	value := s.logic.GetCache(ctx, in.GetUserID())
	return &user.GetUserResponse{UserName: value}, nil
}
