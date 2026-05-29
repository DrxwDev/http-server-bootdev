package chirpy

import (
	"context"
	"errors"

	"github.com/DrxwDev/http-server/internal/database"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, body, userID string) (Chirp, error)
	FindAll(ctx context.Context) ([]Chirp, error)
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

func (r *chirpRepository) FindAll(ctx context.Context) ([]Chirp, error) {
	list, err := r.q.GetAllChirps(ctx)
	if err != nil {
		return nil, errors.New("error getting all chirps")
	}

	chirpList := make([]Chirp, len(list))

	for i, c := range list {
		chirp := chirpCreated(c)
		chirpList[i] = chirp
	}

	return chirpList, nil
}
