package main

import (
	"log"

	"github.com/DrxwDev/http-server/internal/bootdev"
	"github.com/DrxwDev/http-server/internal/config"
	"github.com/DrxwDev/http-server/internal/database"
	"github.com/DrxwDev/http-server/internal/logger"
	"github.com/DrxwDev/http-server/internal/middlewares"
	"github.com/DrxwDev/http-server/internal/server"
	"github.com/DrxwDev/http-server/internal/users"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file")
	}
}

func main() {
	fx.New(
		config.Module,
		logger.Module,
		server.Module,
		bootdev.Module,
		middlewares.Module,
		database.Module,
		users.Module,
	).Run()
}
