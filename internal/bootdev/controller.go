package bootdev

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BootController struct{}

func NewBootController() *BootController {
	return &BootController{}
}

func (c *BootController) Reset(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (c *BootController) App(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to Chirpy")
}

func (c *BootController) Assets(ctx *gin.Context) {
	ctx.String(http.StatusOK, "<a href=\"logo.png\">logo.png</a>")
}

func (c *BootController) Metrics(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (c *BootController) Healthz(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func (c *BootController) ValidateChirp(ctx *gin.Context) {
	var request RequestDTO

	if err := ctx.ShouldBindJSON(&request); err != nil || request.Length() == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	if request.Length() > 140 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Chirp is too long",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"valid": true,
	})
}
