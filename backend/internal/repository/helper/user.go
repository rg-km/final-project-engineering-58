package helper

import (
	"context"
	"backend/domain/entity"
	"backend/internal/repository/mapper"
	"backend/internal/repository/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func StoreUser(ctx context.Context, E dbq.EFn, payload *entity.User) error {
	userDbq := mapper.ToDbqStructUser(payload)
	stmt := dbq.INSERTStmt(models.User{}.TableName(), models.TableUsers(), len(userDbq), dbq.MySQL)

	_, err := E(ctx, stmt, nil, userDbq)

	if err != nil {
		return err
	}

	return nil
}