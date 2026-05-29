package chirpy

import (
	"context"

	"github.com/DrxwDev/http-server/internal/database"
	"github.com/google/uuid"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, body, userID string) (Chirp, error)
	FindAll(ctx context.Context) ([]Chirp, error)
	FindByID(ctx context.Context, id string) (Chirp, error)
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
		return nil, ErrChirpsListNotFound
	}

	chirpList := make([]Chirp, len(list))

	for i, c := range list {
		chirp := chirpCreated(c)
		chirpList[i] = chirp
	}

	return chirpList, nil
}

func (r *chirpRepository) FindByID(ctx context.Context, id string) (Chirp, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return Chirp{}, ErrChirpIDNotParsed
	}

	chirp, err := r.q.FindChirp(ctx, uid)
	if err != nil {
		return Chirp{}, ErrChirpNotFound
	}

	find := chirpCreated(chirp)

	return find, nil
}
