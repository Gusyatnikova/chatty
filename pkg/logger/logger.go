package logger

import (
	"chatty/chatty/usecase"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() usecase.ChatLogger {
	zl, _ := zap.NewDevelopment()
	defer zl.Sync()

	return &Logger{logger: zl}
}

func (e *Logger) Debug(msg string) {
	e.logger.Debug(msg)
}

func (e *Logger) Info(msg string) {
	e.logger.Info(msg)
}

func (e *Logger) Warn(msg string) {
	e.logger.Warn(msg)
}

func (e *Logger) Error(msg string) {
	e.logger.Error(msg)
}

func (e *Logger) Fatal(msg string) {
	e.logger.Fatal(msg)
}

func (e *Logger) Panic(msg string) {
	e.logger.Panic(msg)
}
