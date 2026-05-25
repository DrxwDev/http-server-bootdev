package server

import (
	"github.com/DrxwDev/http-server/internal/bootdev"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, bootdev *bootdev.BootController) {
	router.GET("/", bootdev.Index)
	router.GET("/assets/logo.png", bootdev.Assets)
}
