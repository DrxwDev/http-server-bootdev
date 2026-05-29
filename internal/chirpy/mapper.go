package chirpy

import (
	"github.com/DrxwDev/http-server/internal/database"
	"github.com/google/uuid"
)

func chirpParams(body, userID string) database.CreateChirpParams {
	uid, _ := uuid.Parse(userID)

	chirpParams := database.CreateChirpParams{
		ID:     uuid.New(),
		Body:   body,
		UserID: uid,
	}

	return chirpParams
}

func chirpCreated(chirp database.Chirp) Chirp {
	return Chirp{
		ID:        chirp.ID.String(),
		CreatedAt: chirp.CreatedAt.String(),
		UpdatedAt: chirp.UpdatedAt.String(),
		Body:      chirp.Body,
		UserID:    chirp.UserID.String(),
	}
}

func chirpDTO(chirp Chirp) ChirpResponseDTO {
	return ChirpResponseDTO{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}
}
