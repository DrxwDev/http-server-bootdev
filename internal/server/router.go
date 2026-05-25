package server

import "github.com/gin-gonic/gin"

func NewGinRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.HandleMethodNotAllowed = true

	return router
}
