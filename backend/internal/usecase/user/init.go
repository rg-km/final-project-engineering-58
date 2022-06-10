package user

import (
	"backend/domain/repository"
	"backend/internal/config"
)

type userInteractor struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

func NewUserInteractor(repo repository.UserRepository, config *config.Config) *userInteractor {

	return &userInteractor{
		userRepo: repo,
		cfg:      config,
	}
}

func LoginInteractor(repo repository.UserRepository, config *config.Config) *userInteractor {

	return &userInteractor{
		userRepo: repo,
		cfg:      config,
	}
}
