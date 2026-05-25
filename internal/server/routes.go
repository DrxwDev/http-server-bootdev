package server

import (
	"github.com/DrxwDev/http-server/internal/bootdev"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, bootdev *bootdev.BootController) {
	router.GET("/app", bootdev.Index)
	router.GET("/app/assets", bootdev.Assets)
	router.GET("/healthz", bootdev.Health)
}
