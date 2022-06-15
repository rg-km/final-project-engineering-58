package category

import (
	"backend/domain/repository"
	"backend/internal/config"
)

type categoryInteractor struct {
	repo repository.CategoryRepository
	cfg      *config.Config
}

func NewCategoryInteractor(repo repository.CategoryRepository, config *config.Config) *categoryInteractor {
	return &categoryInteractor{
		repo: repo,
		cfg: config,
	}
}