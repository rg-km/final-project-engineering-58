package user

import (
	"backend/domain/entity"
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *userInteractor) Update(ctx context.Context, payload entity.UserDto, id string) (*entity.User, *exceptions.CustomError) {
	var multilerr *multierror.Error

	user, errFind := x.userRepo.FindByID(ctx, id)
	if errFind != nil {
		multilerr = multierror.Append(multilerr, errFind)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	update := entity.NewUser(payload)

	update.SetPasswordHash(payload.Password)

	update.SetRole("member")

	if err := update.Validate(); err != nil {
		return nil, &exceptions.CustomError{
			Errors: err,
			Status: exceptions.ERRDOMAIN,
		}
	}

	err := x.userRepo.Update(ctx, update, id)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	user.Name = update.Name
	user.Email = update.Email
	user.Password = update.Password

	return user, nil
}
