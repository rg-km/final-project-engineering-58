// package login

// import (
// 	"backend/domain/entity"
// 	"backend/pkg/exceptions"
// 	"context"
// 	"errors"

// 	"github.com/hashicorp/go-multierror"
// )

// func (x *userInteractor) Create(ctx context.Context, payload entity.UserDto) (*entity.User, *exceptions.CustomError) {
// 	var multilerr *multierror.Error

// 	create := entity.NewUser(payload)

// 	create.SetRole("member")

// 	errFeedback := create.Validate()

// 	if errFeedback != nil {
// 		return nil, &exceptions.CustomError{
// 			Errors: errFeedback,
// 			Status: exceptions.ERRDOMAIN,
// 		}
// 	}

// 	findEmail, _ := x.userRepo.FindByEmail(ctx, create.Email)
// 	if findEmail != nil {
// 		multilerr = multierror.Append(multilerr, errors.New("email already exist"))
// 		return nil, &exceptions.CustomError{
// 			Status: exceptions.ERRREPOSITORY,
// 			Errors: multilerr,
// 		}
// 	}

// 	errCreate := x.userRepo.Create(ctx, create)
// 	if errCreate != nil {
// 		multilerr = multierror.Append(multilerr, errCreate)
// 		return nil, &exceptions.CustomError{
// 			Errors: multilerr,
// 			Status: exceptions.ERRREPOSITORY,
// 		}
// 	}

// 	return create, nil
// }
