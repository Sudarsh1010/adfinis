package di

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	// Set defaults for logger config
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_FORMAT", "console")

	// Bind to environment variables
	viper.AutomaticEnv()

	// Parse log level
	var zapLevel zapcore.Level
	switch viper.GetString("LOG_LEVEL") {
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

	// Build zap config based on format
	var cfgZap zap.Config
	if viper.GetString("LOG_FORMAT") == "json" {
		cfgZap = zap.NewProductionConfig()
	} else {
		cfgZap = zap.NewDevelopmentConfig()
	}
	cfgZap.Level = zap.NewAtomicLevelAt(zapLevel)

	return cfgZap.Build()
}
