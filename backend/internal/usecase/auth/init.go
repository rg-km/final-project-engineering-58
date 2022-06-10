package auth

import (
	"backend/domain/repository"
	"backend/internal/config"
)

type authInteractor struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

func NewAuthInteractor(repo repository.UserRepository, config *config.Config) *authInteractor {

	return &authInteractor{
		userRepo: repo,
		cfg:      config,
	}
}