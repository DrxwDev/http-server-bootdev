package bootdev

import (
	"context"
	"strings"
)

type BootService interface {
	CleanChirpy(ctx context.Context, chirpy string) (string, bool)
}

type bootService struct{}

func NewBootService() BootService {
	return &bootService{}
}

func (s *bootService) CleanChirpy(ctx context.Context, chirpy string) (string, bool) {
	find := false
	profane := [3]string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Fields(chirpy)

	for i, w := range words {
		for _, bad := range profane {
			if strings.ToLower(w) == bad {
				words[i] = "****"
				find = true
			}
		}
	}

	return strings.Join(words, " "), find
}
