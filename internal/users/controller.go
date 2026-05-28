package users

import (
	"net/http"

	"github.com/DrxwDev/http-server/internal/config"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service UserService
	cfg     config.AppConfig
}

func NewUserController(service UserService, cfg config.AppConfig) UserController {
	return UserController{
		service: service,
		cfg:     cfg,
	}
}

func (c UserController) CreateUser(ctx *gin.Context) {
	var request UserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.CreateUser(ctx.Request.Context(), request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := userDTO(user)

	ctx.JSON(http.StatusCreated, dto)
}

func (c UserController) ResetUsers(ctx *gin.Context) {
	if c.cfg.PLATFORM != "dev" {
		ctx.Status(http.StatusForbidden)
		return
	}

	err := c.service.ResetUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "users table restarted",
	})
}
