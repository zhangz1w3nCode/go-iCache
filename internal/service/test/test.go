package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	"github.com/zhangz1w3nCode/go-iCache/internal/logic/cache"
	"log"
)

type TestService struct {
	helloworld.UnimplementedTestServiceServer
	logic *cache.TestLogic
}

func NewTestService() *TestService {
	return &TestService{
		logic: cache.NewTestLogic(),
	}
}

func (s *TestService) TestConnect(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	s.logic.SetCache(ctx, in.GetName(), "testSet")
	value := s.logic.GetCache(ctx, in.GetName())

	return &helloworld.HelloReply{Message: value}, nil
}
