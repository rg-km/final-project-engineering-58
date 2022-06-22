package repository

import (
	"backend/domain/entity"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, payload *entity.Post) error
	FindByID(ctx context.Context, id string) (*entity.Post, error)
	GetAll(ctx context.Context) ([]*entity.Post, error)
	Update(ctx context.Context, payload *entity.Post, id string) error
	Delete(ctx context.Context, id string) error
}