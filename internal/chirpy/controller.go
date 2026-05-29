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

func (c ChirpController) GetAll(ctx *gin.Context) {
	list, err := c.service.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chirpList := make([]ChirpResponseDTO, len(list))

	for i, c := range list {
		chirp := chirpDTO(c)
		chirpList[i] = chirp
	}

	ctx.JSON(http.StatusOK, chirpList)
}

func (c ChirpController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	chirp, err := c.service.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, chirpDTO(chirp))
}
