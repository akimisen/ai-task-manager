package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, _ := config.Build()
	sugar := zapLogger.Sugar()

	return &Logger{
		SugaredLogger: sugar,
	}
}

// 确保在程序退出时同步日志
func (l *Logger) Sync() error {
	return l.SugaredLogger.Sync()
}
