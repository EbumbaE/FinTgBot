package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {

	debugMode := os.Getenv("DEBUG_MODE")

	var localLogger *zap.Logger
	var err error

	switch debugMode {
	case "prod":
		localLogger, err = zap.NewProduction()
	default:
		cfg := zap.NewDevelopmentConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		localLogger, err = cfg.Build()
	}
	if err != nil {
		log.Fatal("logger init: ", err)
	}

	logger = localLogger
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
