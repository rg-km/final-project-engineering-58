package usecase

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"
)

type UserUsecase interface {
	Create(ctx context.Context, payload entity.UserDto) (*entity.User, *exceptions.CustomError)
}