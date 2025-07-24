package logger

import (
	"os"

	"go.uber.org/zap"
)

var ZapLogger *zap.Logger

func init() {
	ZapLogger = zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		ZapLogger = zap.Must(zap.NewDevelopment())
	}
}

func Sync() {
	ZapLogger.Sync()
}

func Info(message string, fields ...zap.Field) {
	ZapLogger.Info(message, fields...)
}


func Debug(message string, fields ...zap.Field) {
	ZapLogger.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	ZapLogger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	ZapLogger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	ZapLogger.Fatal(message, fields...)
}