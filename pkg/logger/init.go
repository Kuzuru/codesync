package logger

import (
	"go.uber.org/zap"
)

func Init() (*zap.Logger, error) {
	logger, err := zap.NewProduction()

	defer logger.Sync()

	return logger, err
}
