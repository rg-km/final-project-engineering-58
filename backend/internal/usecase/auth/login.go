package auth

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"backend/pkg/utils"
	"context"
	"errors"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) Login(ctx context.Context, payload entity.Auth) (*entity.User, *exceptions.CustomError) {
	var multilerr *multierror.Error

	err := payload.Validate()
	
	if err != nil {
		return nil, &exceptions.CustomError{
			Errors: err,
			Status: exceptions.ERRDOMAIN,
		}
	}

	user, errUser := x.userRepo.FindByEmail(ctx, payload.Email)
	if errUser != nil {
		multilerr = multierror.Append(multilerr, errUser)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	errLogin := utils.CheckPasswordHash(payload.Password, user.Password)
	if !errLogin {
		multilerr = multierror.Append(multilerr, errors.New("incorrect password. please enter a password that matches your registered email"))
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	return user, nil
}