package chirpy

import (
	"context"
	"strings"
)

type ChirpService interface {
	CreateChirp(ctx context.Context, body, userID string) (Chirp, error)
	CleanChirpy(ctx context.Context, chirpy string) string
	FindAll(ctx context.Context) ([]Chirp, error)
	FindByID(ctx context.Context, id string) (Chirp, error)
}

type chirpService struct {
	repo ChirpRepository
}

func NewChirpService(repo ChirpRepository) ChirpService {
	return &chirpService{
		repo: repo,
	}
}

func (s *chirpService) CleanChirpy(ctx context.Context, chirpy string) string {
	profane := [3]string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Fields(chirpy)

	for i, w := range words {
		for _, bad := range profane {
			if strings.ToLower(w) == bad {
				words[i] = "****"
			}
		}
	}

	return strings.Join(words, " ")
}

func (s *chirpService) CreateChirp(ctx context.Context, body string, userID string) (Chirp, error) {
	if body == "" {
		return Chirp{}, ErrEmptyBody
	}

	if userID == "" {
		return Chirp{}, ErrEmptyUserID
	}

	if len(body) > 140 {
		return Chirp{}, ErrChirpTooLong
	}

	chirp := s.CleanChirpy(ctx, body)

	return s.repo.CreateChirp(ctx, chirp, userID)
}

func (s *chirpService) FindAll(ctx context.Context) ([]Chirp, error) {
	return s.repo.FindAll(ctx)
}

func (s *chirpService) FindByID(ctx context.Context, id string) (Chirp, error) {
	if id == "" {
		return Chirp{}, ErrChirpIDNotFound
	}

	return s.repo.FindByID(ctx, id)
}
