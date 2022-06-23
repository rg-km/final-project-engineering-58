package repository

import (
	"backend/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, payload *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, payload *entity.User, id string) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*entity.User, error)
}
