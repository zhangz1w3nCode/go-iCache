package test

import (
	"context"
	"github.com/zhangz1w3nCode/go-iCache/internal/api/generate/helloworld"
	"log"
)

type TestService struct {
	helloworld.UnimplementedTestServiceServer
}

func NewTestService() *TestService {
	return &TestService{}
}

func (s *TestService) TestConnect(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
