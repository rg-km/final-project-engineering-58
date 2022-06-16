package repository

import (
	"backend/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, payload *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	// Update(ctx context.Context, payload *entity.User, id string) error
}
