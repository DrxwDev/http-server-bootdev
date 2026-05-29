package chirpy

import (
	"context"

	"github.com/DrxwDev/http-server/internal/database"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, body, userID string) (Chirp, error)
}

type chirpRepository struct {
	q *database.Queries
}

func NewChirpRepository(q *database.Queries) ChirpRepository {
	return &chirpRepository{
		q: q,
	}
}

func (r *chirpRepository) CreateChirp(ctx context.Context, body string, userID string) (Chirp, error) {
	params := chirpParams(body, userID)

	chirp, err := r.q.CreateChirp(ctx, params)
	if err != nil {
		return Chirp{}, ErrChirpNotCreated
	}

	created := chirpCreated(chirp)
	return created, nil
}
