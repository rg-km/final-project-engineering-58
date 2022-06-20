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

func (x *UserRepository) Update(ctx context.Context, payload *entity.User, id string) error {
	fields := map[string]interface{}{
		"name":     payload.Name,
		"email":    payload.Email,
		"password": payload.Password,
	}

	stmt := builder.UPDATEStmt(models.User{}.TableName(), fields, "id", id)

	_, err := builder.QueryExec(ctx, x.db, stmt, nil)

	if err != nil {
		return err
	}

	return nil
}

func (x *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = ? `, models.User{}.TableName())
	opts := &dbq.Options{
		ConcreteStruct: models.User{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
		SingleResult:   true,
	}

	result := dbq.MustQ(ctx, x.db, stmt, opts, id)

	if result == nil {
		return nil, errors.New("user not found")
	}
	user := mapper.ToDomainUser(result.(*models.User))
	return user, nil
}

func (x *UserRepository) Delete(ctx context.Context, id string) error {
	stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, models.User{}.TableName())
	_, err := builder.QueryExec(ctx, x.db, stmt, id)

	if err != nil {
		return err
	}

	return nil
}

func (x *UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.User{}.TableName())
	opts := &dbq.Options{
		ConcreteStruct: models.User{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result := dbq.MustQ(ctx, x.db, stmt, opts)

	if result == nil {
		return nil, errors.New("user not found")
	}
	users := mapper.ToListDomainUser(result.([]*models.User))
	return users, nil
}
