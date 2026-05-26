package server

import (
	"github.com/DrxwDev/http-server/internal/bootdev"
	"github.com/DrxwDev/http-server/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, bootdev *bootdev.BootController, apiConfig *middlewares.APIConfig) {
	app := router.Group("/app", middlewares.CountMiddleware(apiConfig))
	admin := router.Group("/admin")

	app.GET("/", bootdev.App)

	admin.POST("/reset", middlewares.ResetMiddleware(apiConfig), bootdev.Reset)
	admin.GET("/metrics", middlewares.MetricsMiddleware(apiConfig))
}
