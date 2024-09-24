package cacheWarn

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func SendWarnMessage(warnMsg string) {
	//初始化日志 logger
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	//todo:发送告警信息
	logger.Warn(warnMsg)
}
