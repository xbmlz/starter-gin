package logger

import "go.uber.org/zap"

func NewZapLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
