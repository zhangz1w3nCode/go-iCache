// Package service 是 GRPC 服务的实现，需要通过 wire 进行依赖注入来初始化所有的服务
package service

import "visual-state-machine/internal/service/voice"

// Services 是所有服务的集合
type Services struct {
	HookService *voice.HookService
}

func newServices(
	hookService *voice.HookService) *Services {
	return &Services{
		HookService: hookService,
	}
}
