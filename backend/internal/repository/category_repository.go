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

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return CategoryRepository{db: db}
}

func (x *CategoryRepository) Create(ctx context.Context, payload *entity.Category) error {
	args := mapper.ToDbqStructCategory(payload)
	stmt := builder.INSERTStmt(models.Category{}.TableName(), models.TableCategories(), len(args))

	_, err := builder.QueryExec(ctx, x.db, stmt, args)

	if err != nil {
		return err
	}

	return nil
}

func (x *CategoryRepository) FindByID(ctx context.Context, id string) (*entity.Category, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = ? `, models.Category{}.TableName())
	opts := &dbq.Options{
		ConcreteStruct: models.Category{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
		SingleResult:   true,
	}

	result := dbq.MustQ(ctx, x.db, stmt, opts, id)

	if result == nil {
		return nil, errors.New("category not found")
	}
	category := mapper.ToDomainCategory(result.(*models.Category))
	return category, nil
}

func (x *CategoryRepository) GetAll(ctx context.Context) ([]*entity.Category, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.Category{}.TableName())
	opts := &dbq.Options{
		ConcreteStruct: models.Category{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
		SingleResult:   false,
	}

	result := dbq.MustQ(ctx, x.db, stmt, opts)

	if result == nil {
		return nil, errors.New("category not found")
	}

	categories := mapper.ToListDomainCategory(result.([]*models.Category))
	return categories, nil
}

func (x *CategoryRepository) Update(ctx context.Context, payload *entity.Category, id string) error {
	fields := map[string]interface{}{
		"name": payload.Name,
	}

	stmt := builder.UPDATEStmt(models.Category{}.TableName(), fields, "id", id)

	_, err := builder.QueryExec(ctx, x.db, stmt, nil)

	if err != nil {
		return err
	}

	return nil
}

func (x *CategoryRepository) Delete(ctx context.Context, id string) error {
	stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, models.Category{}.TableName())
	_, err := builder.QueryExec(ctx, x.db, stmt, id)

	if err != nil {
		return err
	}

	return nil
}