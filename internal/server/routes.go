package server

import (
	"github.com/DrxwDev/http-server/internal/bootdev"
	"github.com/DrxwDev/http-server/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, bootdev *bootdev.BootController, apiConfig *middlewares.APIConfig) {
	router.GET("/app", middlewares.CountMiddleware(apiConfig), bootdev.App)
	router.GET("/metrics", middlewares.MetricsMiddleware(apiConfig), bootdev.Metrics)
	router.GET("/healthz", bootdev.Healthz)
	router.POST("/reset", middlewares.ResetMiddleware(apiConfig), bootdev.Reset)
}
