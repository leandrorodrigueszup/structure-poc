package tracking

import (
	"context"
	"go.uber.org/zap"
	"poc/internal/configuration"
)

type LoggerType string

const LoggerFlag = LoggerType("api-logger-context")

func NewLogger() (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if configuration.IsRunningInProduction() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	return ctx.Value(LoggerFlag).(*zap.SugaredLogger)
}
