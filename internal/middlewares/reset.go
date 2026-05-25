package middlewares

import "github.com/gin-gonic/gin"

func ResetMiddleware(api *APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		api.fileserverHits.Store(0)
		ctx.Next()
	}
}
