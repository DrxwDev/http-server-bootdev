package users

import (
	"github.com/DrxwDev/http-server/internal/database"
	"github.com/google/uuid"
)

func userParams(email string) database.CreateUserParams {
	user := database.CreateUserParams{
		ID:    uuid.New(),
		Email: email,
	}

	return user
}

func userCreated(u database.User) User {
	user := User{
		ID:        u.ID.String(),
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
		Email:     u.Email,
	}

	return user
}

func userDTO(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
}
