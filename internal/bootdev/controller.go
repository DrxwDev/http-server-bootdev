package bootdev

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BootController struct{}

func NewBootController() *BootController {
	return &BootController{}
}

func (c *BootController) Index(ctx *gin.Context) {
	ctx.File("assets/index.html")
	ctx.Status(http.StatusOK)
}

func (c *BootController) Assets(ctx *gin.Context) {
	ctx.Header("Content-Type", "image/png")
	ctx.File("assets/logo.png")
	ctx.Status(http.StatusOK)
}
