package usecase

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"
)

type CategoryUsecase interface {
	Create(ctx context.Context, payload entity.CategoryDto) (*entity.Category, *exceptions.CustomError)
	GetAll(ctx context.Context) ([]*entity.Category, *exceptions.CustomError)
	Update(ctx context.Context, payload entity.CategoryDto, id string) (*entity.Category, *exceptions.CustomError)
	Delete(ctx context.Context, id string) *exceptions.CustomError
}
