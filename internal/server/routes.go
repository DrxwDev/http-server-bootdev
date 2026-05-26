package server

import (
	"github.com/DrxwDev/http-server/internal/bootdev"
	"github.com/DrxwDev/http-server/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, bootdev *bootdev.BootController, apiConfig *middlewares.APIConfig) {
	app := router.Group("/app", middlewares.CountMiddleware(apiConfig))
	api := router.Group("/api")

	app.GET("/", bootdev.App)

	api.POST("/validate_chirp", bootdev.ValidateChirp)
}
