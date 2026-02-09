package di

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	ReadTimeout = 15 * time.Second
	WriteTimout = 15 * time.Second
	IdleTimeout = 60 * time.Second

	BunMaxOpenConns    = 25
	BunMaxIdleConns    = 5
	BunConnMaxLifetime = 5 * time.Minute
	BunConnMaxIdleTime = 5 * time.Minute

	HTTPDefaultPort = 3000
	HTTPMaxAge      = 300
)

type Config struct {
	Port      int
	Env       string
	DBPath    string
	LogLevel  string
	LogFormat string
}

func NewConfig(logger *Logger) (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		logger.Infoln("No .env file found, using environment variables")
	}

	viper.SetDefault("PORT", HTTPDefaultPort)
	viper.SetDefault("ENV", "development")
	viper.SetDefault("DB_PATH", "./adfinis.db")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_FORMAT", "console")

	viper.AutomaticEnv()

	port := viper.GetInt("PORT")
	if port == 0 {
		port = 3000
	}

	return &Config{
		Port:      port,
		Env:       viper.GetString("ENV"),
		DBPath:    viper.GetString("DB_PATH"),
		LogLevel:  viper.GetString("LOG_LEVEL"),
		LogFormat: viper.GetString("LOG_FORMAT"),
	}, nil
}
