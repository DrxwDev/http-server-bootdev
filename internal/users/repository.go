package users

import (
	"context"

	"github.com/DrxwDev/http-server/internal/database"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string) (User, error)
	ResetUsers(ctx context.Context) error
}

type userRepository struct {
	q *database.Queries
}

func NewUserRepository(queries *database.Queries) UserRepository {
	return &userRepository{
		q: queries,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, email string) (User, error) {
	params := userParams(email)
	user, err := r.q.CreateUser(ctx, params)
	if err != nil {
		return User{}, err
	}

	newUser := userCreated(user)
	return newUser, nil
}

func (r *userRepository) ResetUsers(ctx context.Context) error {
	return r.q.ResetUsers(ctx)
}
