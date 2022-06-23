package auth

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) Profile(ctx context.Context, userId string) (*entity.User, *exceptions.CustomError) {
	var multilerr *multierror.Error

	user, err := x.userRepo.FindByID(ctx, userId)

	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return user, nil
}
