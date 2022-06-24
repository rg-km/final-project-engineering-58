package user

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *userInteractor) GetAll(ctx context.Context) ([]*entity.User, *exceptions.CustomError) {
	var multilerr *multierror.Error

	users, err := x.userRepo.GetAll(ctx)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	return users, nil
}
