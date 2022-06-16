package usecase

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"
)

type AuthUsecase interface {
	Login(context context.Context, payload entity.Auth) (*entity.User, *exceptions.CustomError)
	Register(context context.Context, payload entity.UserDto) (*entity.User, *exceptions.CustomError)
}