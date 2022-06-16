package repository

import (
	"context"
	"backend/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, payload *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindById(ctx context.Context, id string) (*entity.User, error)
}