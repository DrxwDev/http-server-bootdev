package middlewares

import "github.com/gin-gonic/gin"

func CountMiddleware(api *APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		api.fileserverHits.Add(1)
		ctx.Next()
	}
}
