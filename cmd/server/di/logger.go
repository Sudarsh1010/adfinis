package di

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(cfg *Config) (*Logger, error) {
	var zapLevel zapcore.Level
	switch cfg.LogLevel {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	default:
		zapLevel = zap.InfoLevel
	}

	var cfgZap zap.Config
	if cfg.LogFormat == "json" {
		cfgZap = zap.NewProductionConfig()
	} else {
		cfgZap = zap.NewDevelopmentConfig()
	}
	cfgZap.Level = zap.NewAtomicLevelAt(zapLevel)

	z, err := cfgZap.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{z.Sugar()}, nil
}

func (l *Logger) Debug(msg string, fields ...any) {
	l.With(fields...).Debug(msg)
}

func (l *Logger) Info(msg string, fields ...any) {
	l.With(fields...).Info(msg)
}

func (l *Logger) Warn(msg string, fields ...any) {
	l.With(fields...).Warn(msg)
}

func (l *Logger) Error(msg string, fields ...any) {
	l.With(fields...).Error(msg)
}
