package server

import (
	"github.com/DrxwDev/http-server/internal/bootdev"
	"github.com/DrxwDev/http-server/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, bootdev *bootdev.BootController, apiConfig *middlewares.APIConfig) {
	app := router.Group("/app", middlewares.CountMiddleware(apiConfig))
	api := router.Group("/api")

	app.GET("", bootdev.App)
	app.GET("/assets", bootdev.Assets)

	api.GET("/metrics", middlewares.MetricsMiddleware(apiConfig), bootdev.Metrics)
	api.GET("/healthz", bootdev.Healthz)
	api.POST("/reset", middlewares.ResetMiddleware(apiConfig), bootdev.Reset)
}
