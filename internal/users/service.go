package users

import (
	"context"
	"strings"
)

type UserService interface {
	CreateUser(ctx context.Context, email string) (User, error)
	ResetUsers(ctx context.Context) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repo: repository,
	}
}

func (s *userService) CreateUser(ctx context.Context, email string) (User, error) {
	if email == "" {
		return User{}, ErrEmptyEmail
	}

	if !strings.Contains(email, "@") {
		return User{}, ErrInvalidEmail
	}

	return s.repo.CreateUser(ctx, email)
}

func (s *userService) ResetUsers(ctx context.Context) error {
	return s.repo.ResetUsers(ctx)
}
