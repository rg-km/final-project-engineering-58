package category

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *categoryInteractor) GetAll(ctx context.Context) ([]*entity.Category, *exceptions.CustomError) {
	var multilerr *multierror.Error

	categories, err := x.repo.GetAll(ctx)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	return categories, nil
}