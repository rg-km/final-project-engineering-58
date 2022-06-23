package usecase

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"
)

type UserUsecase interface {
	Create(ctx context.Context, payload entity.UserDto) (*entity.User, *exceptions.CustomError)
	//FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, payload entity.UserDto, id string) (*entity.User, *exceptions.CustomError)
	Delete(ctx context.Context, id string) *exceptions.CustomError
	GetAll(ctx context.Context) ([]*entity.User, *exceptions.CustomError)
}
