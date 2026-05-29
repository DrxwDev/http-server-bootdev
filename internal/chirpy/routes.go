package chirpy

import (
	"github.com/gin-gonic/gin"
)

func ChirpRoutes(router *gin.Engine, controller ChirpController) {
	routes := router.Group("/api/v1/chirp")

	routes.POST("", controller.CreateChirp)
	routes.GET("", controller.GetAll)
}
