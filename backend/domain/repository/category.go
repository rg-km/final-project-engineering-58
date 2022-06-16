package repository

import (
	"backend/domain/entity"
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, payload *entity.Category) error
	FindByID(ctx context.Context, id string) (*entity.Category, error)
	GetAll(ctx context.Context) ([]*entity.Category, error)
	Update(ctx context.Context, payload *entity.Category, id string) error
	Delete(ctx context.Context, id string) error
}