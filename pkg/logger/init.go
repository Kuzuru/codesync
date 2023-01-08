package logger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	if !fiber.IsChild() {
		l.Logger.Info(msg, fields...)
	}
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	if !fiber.IsChild() {
		l.Logger.Warn(msg, fields...)
	}
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	if !fiber.IsChild() {
		l.Logger.Error(msg, fields...)
	}
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	if !fiber.IsChild() {
		l.Logger.Panic(msg, fields...)
	}
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	if !fiber.IsChild() {
		l.Logger.Fatal(msg, fields...)
	}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if !fiber.IsChild() {
		l.Logger.Debug(msg, fields...)
	}
}

func New() *Logger {
	logger, _ := zap.NewProduction()

	return &Logger{
		logger,
	}
}
