package users

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine, controller UserController) {
	routes := router.Group("/api/v1/users")

	routes.POST("", controller.CreateUser)
	router.POST("/admin/reset", controller.ResetUsers)
}
