package chirpy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChirpController struct {
	service ChirpService
}

func NewChirpController(service ChirpService) ChirpController {
	return ChirpController{
		service: service,
	}
}

func (c ChirpController) CreateChirp(ctx *gin.Context) {
	var request ChirpRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chirp, err := c.service.CreateChirp(ctx.Request.Context(), request.Body, request.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := chirpDTO(chirp)
	ctx.JSON(http.StatusCreated, dto)
}
