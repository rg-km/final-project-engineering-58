package post

import (
	"backend/domain/repository"
	"backend/internal/config"
)

type postInteractor struct {
	repo repository.PostRepository
	cfg      *config.Config
}

func NewPostInteractor(repo repository.PostRepository, config *config.Config) *postInteractor {
	return &postInteractor{
		repo: repo,
		cfg: config,
	}
}