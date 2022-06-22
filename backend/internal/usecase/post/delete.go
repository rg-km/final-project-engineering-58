package post

import (
	"backend/pkg/exceptions"
	"context"

	"github.com/hashicorp/go-multierror"
)

func (x *postInteractor) Delete(ctx context.Context, id string) *exceptions.CustomError {
	var multilerr *multierror.Error

	_, errFind := x.repo.FindByID(ctx, id)
	if errFind != nil {
		multilerr = multierror.Append(multilerr, errFind)
		return &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}
	
	err := x.repo.Delete(ctx, id)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Errors: multilerr,
			Status: exceptions.ERRREPOSITORY,
		}
	}

	return nil
}