package config

import (
	"os"
	"time"
)

func LoadApp() AppConfig {
	return AppConfig{
		HOST: os.Getenv("HOST"),
		PORT: os.Getenv("PORT"),
	}
}

func LoadServer() ServerConfig {
	return ServerConfig{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}

func LoadDatabase() DatabaseConfig {
	return DatabaseConfig{
		DBDsn: os.Getenv("DB_DSN"),
	}
}
