package repository

import (
	"backend/domain/entity"
	"backend/internal/repository/helper"
	"backend/internal/repository/mapper"
	"backend/internal/repository/models"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rocketlaunchr/dbq/v2"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db: db}
}

func (x *UserRepository) Create(ctx context.Context, payload *entity.User) error {
	var err error
	_ = dbq.Tx(ctx, x.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		err = helper.StoreUser(ctx, E, payload)
		if err != nil {
			return
		}
		_ = txCommit()
		return
	})

	if err != nil {
		return err
	}

	return nil
}

func (x *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ? `, models.User{}.TableName())
	opts := &dbq.Options{
		ConcreteStruct: models.User{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
		SingleResult:   true,
	}

	result := dbq.MustQ(ctx, x.db, stmt, opts, email)

	if result == nil {
		return nil, errors.New("user not found")
	}
	user := mapper.ToDomainUser(result.(*models.User))
	return user, nil
}