package post

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *postInteractor) Update(ctx context.Context, payload entity.PostDto, id string) (*entity.Post, *exceptions.CustomError) {
	var multilerr *multierror.Error

	post, errFind := x.repo.FindByID(ctx, id)
	if errFind != nil {
		multilerr = multierror.Append(multilerr, errFind)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	update := entity.NewPost(payload)

	if err := update.Validate(); err != nil {
		return nil, &exceptions.CustomError{
			Errors: err,
			Status: exceptions.ERRDOMAIN,
		}
	}

	err := x.repo.Update(ctx, update, id)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	post.Title = update.Title
	post.Content = update.Content
	post.UrlVideo = update.UrlVideo
	post.CategoryID = update.CategoryID

	return post, nil
}