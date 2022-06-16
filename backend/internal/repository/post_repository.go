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

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return PostRepository{db: db}
}

func (x *PostRepository) Create(ctx context.Context, payload *entity.Post) error {
	args := mapper.ToDbqStructPost(payload)
	stmt := builder.INSERTStmt(models.Post{}.TableName(), models.TablePosts(), len(args))

	fmt.Println(stmt)
	fmt.Println(args)

	_, err := builder.QueryExec(ctx, x.db, stmt, args)

	if err != nil {
		return err
	}

	return nil
}

func (x *PostRepository) FindByID(ctx context.Context, id string) (*entity.Post, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = ? `, models.Post{}.TableName())
	opts := &dbq.Options{
		ConcreteStruct: models.Post{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
		SingleResult:   true,
	}

	result := dbq.MustQ(ctx, x.db, stmt, opts, id)

	if result == nil {
		return nil, errors.New("post not found")
	}
	post := mapper.ToDomainPost(result.(*models.Post))
	return post, nil
}

func (x *PostRepository) GetAll(ctx context.Context) ([]*entity.Post, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.Post{}.TableName())
	opts := &dbq.Options{
		ConcreteStruct: models.Post{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
		SingleResult:   false,
	}

	result := dbq.MustQ(ctx, x.db, stmt, opts)

	if result == nil {
		return nil, errors.New("post not found")
	}

	categories := mapper.ToListDomainPost(result.([]*models.Post))
	return categories, nil
}

func (x *PostRepository) Update(ctx context.Context, payload *entity.Post, id string) error {
	fields := map[string]interface{}{
		"title":        payload.Title,
		"description":  payload.Description,
		"content":      payload.Content,
		"url_video":    payload.UrlVideo,
		"category_id":  payload.CategoryID,
	}

	stmt := builder.UPDATEStmt(models.Post{}.TableName(), fields, "id", id)

	_, err := builder.QueryExec(ctx, x.db, stmt, nil)

	if err != nil {
		return err
	}

	return nil
}

func (x *PostRepository) Delete(ctx context.Context, id string) error {
	stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, models.Post{}.TableName())
	_, err := builder.QueryExec(ctx, x.db, stmt, id)

	if err != nil {
		return err
	}

	return nil
}