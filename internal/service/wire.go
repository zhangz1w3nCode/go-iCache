//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package service

import (
	"github.com/google/wire"
	voicesvc "visual-state-machine/internal/service/voice"
)

// InitServices 初始化所有服务
func InitServices() *Services {
	wire.Build(newServices,
		voicesvc.NewHookService)
	return &Services{}
}
