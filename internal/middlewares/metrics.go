package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(api *APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hits: %d", api.fileserverHits.Load())
		ctx.Next()
	}
}
