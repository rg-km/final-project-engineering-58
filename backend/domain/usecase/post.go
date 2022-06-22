package usecase

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"
)

type PostUsecase interface {
	Create(ctx context.Context, payload entity.PostDto) (*entity.Post, *exceptions.CustomError)
	GetAll(ctx context.Context) ([]*entity.Post, *exceptions.CustomError)
	Update(ctx context.Context, payload entity.PostDto, id string) (*entity.Post, *exceptions.CustomError)
	Delete(ctx context.Context, id string) *exceptions.CustomError
}
