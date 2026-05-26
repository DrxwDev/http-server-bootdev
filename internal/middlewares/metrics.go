package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(api *APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		count := api.fileserverHits.Load()

		if count <= 0 {
			ctx.String(http.StatusOK, "Welcome, Chirpy Admin\nChirpy has been visited 0 times!")
			ctx.Next()
		}

		ctx.String(http.StatusOK, "Chirpy has been visited %d times!", count)
		ctx.Next()
	}
}
