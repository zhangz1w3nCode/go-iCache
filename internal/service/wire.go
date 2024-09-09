//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package service

import (
	"github.com/google/wire"
	callingsvc "visual-state-machine/internal/service/calling"
	flowsvc "visual-state-machine/internal/service/flow"
)

// InitServices 初始化所有服务
func InitServices() *Services {
	wire.Build(newServices,
		flowsvc.NewFlowService, callingsvc.NewCallingService)
}
