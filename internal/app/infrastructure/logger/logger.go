package logger

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	return cfg.Build()
}
