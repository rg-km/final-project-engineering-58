package repository

import (
	"backend/domain/entity"
	"backend/internal/repository/mapper"
	"backend/internal/repository/models"
	"backend/pkg/builder"
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
	args := mapper.ToDbqStructUser(payload)
	stmt := builder.INSERTStmt(models.User{}.TableName(), models.TableUsers(), len(args))

	_, err := builder.QueryExec(ctx, x.db, stmt, args)

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

// func (x *UserRepository) Update(ctx context.Context, payload *entity.UserDto, id string) error {
// 	fields := map[string]interface{}{
// 		"name": payload.Name,
// 	}

// 	stmt := builder.UPDATEStmt(models.User{}.TableName(), fields, "id", id)

// 	_, err := builder.QueryExec(ctx, x.db, stmt, nil)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
