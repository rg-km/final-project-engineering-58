package post

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *postInteractor) Create(ctx context.Context, payload entity.PostDto) (*entity.Post, *exceptions.CustomError) {
	var multilerr *multierror.Error

	create := entity.NewPost(payload)

	if err := create.Validate(); err != nil {
		return nil, &exceptions.CustomError{
			Errors: err,
			Status: exceptions.ERRDOMAIN,
		}
	}

	err := x.repo.Create(ctx, create)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	return create, nil
}