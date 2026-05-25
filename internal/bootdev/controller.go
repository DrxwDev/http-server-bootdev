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
	ctx.File("assets/assets.html")
	ctx.Status(http.StatusOK)
}

func (c *BootController) Health(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/plain; charset=utf-8")
	ctx.String(http.StatusOK, "OK")
}
