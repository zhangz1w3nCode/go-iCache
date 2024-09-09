// Package service 是 GRPC 服务的实现，需要通过 wire 进行依赖注入来初始化所有的服务
package service

import (
	"visual-state-machine/internal/service/flow"
)

// Services 是所有服务的集合
type Services struct {
	CallingService *flow.ServiceTest
}

func newServices(callingService *flow.ServiceTest) *Services {
	return &Services{
		CallingService: callingService,
	}
}
